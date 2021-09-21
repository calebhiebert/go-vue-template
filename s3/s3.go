package s3

import (
	"bytes"
	"context"
	"fmt"
	"github.com/calebhiebert/go-vue-template/api"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"image"
	"io"
	"net/http"
	"os"
)

var S3 *minio.Client

var MaxImageSize int64 = 5242880

var AllowedMimes = []string{
	"image/jpeg",
	"image/png",
	"image/webm",
}

func InitS3() error {
	minioClient, err := minio.New(os.Getenv("S3_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("S3_KEY"), os.Getenv("S3_SECRET"), ""),
		Secure: os.Getenv("S3_SSL") == "true",
	})
	if err != nil {
		return err
	}

	S3 = minioClient

	err = ensureBucket(context.Background(), os.Getenv("S3_IMAGE_BUCKET_NAME"))
	if err != nil {
		return err
	}

	return nil
}

func UploadImage(c *gin.Context) {
	id := uuid.Must(uuid.NewV4()).String()

	mpf, err := c.MultipartForm()
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	fls, ok := mpf.File["image"]
	if !ok || len(fls) == 0 {
		api.NewAPIError("missing-image", http.StatusBadRequest, "Missing image field").Respond(c)
		return
	}

	ff := fls[0]

	file, err := ff.Open()
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	defer file.Close()

	if mimeOK := mimeAllowed(ff.Header.Get("Content-Type"), AllowedMimes); !mimeOK {
		api.NewAPIError("invalid-image-mime", http.StatusBadRequest, "That filetype is not allowed").
			Respond(c)
		return
	}

	var buf bytes.Buffer

	fileReader := io.TeeReader(file, &buf)

	img, t, err := image.Decode(fileReader)
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	_, err = S3.PutObject(c.Request.Context(), os.Getenv("S3_IMAGE_BUCKET_NAME"), id, &buf, ff.Size, minio.PutObjectOptions{})
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	dbImage := models.Image{
		ID:     id,
		Name:   ff.Filename,
		Type:   ff.Header.Get("Content-Type"),
		Size:   int(ff.Size),
		Width:  img.Bounds().Max.X,
		Height: img.Bounds().Max.Y,
	}

	err = dbImage.InsertG(c.Request.Context(), boil.Infer())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	c.JSON(http.StatusOK, dbImage)
}

func GetImage(c *gin.Context) {
	id := c.Param("id")

	img, err := models.Images(qm.Where("id = ?", id)).OneG(c.Request.Context())
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	obj, err := S3.GetObject(c.Request.Context(), os.Getenv("S3_IMAGE_BUCKET_NAME"), img.ID, minio.GetObjectOptions{})
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	defer obj.Close()

	stat, err := obj.Stat()
	if err != nil {
		api.APIErrorFromErr(err).Respond(c)
		return
	}

	fmt.Println("SIZE", img.Size, stat.Size)

	c.Writer.Header().Add("Content-Type", img.Type)
	c.Writer.Header().Add("Content-Length", fmt.Sprintf("%d", stat.Size))
	c.Writer.WriteHeader(http.StatusOK)

	_, err = io.Copy(c.Writer, obj)
	if err != nil {
		fmt.Println("FILE SEND ERR", err.Error())
		// Can't really do anything at this point
	}
}

func mimeAllowed(mime string, allowed []string) bool {
	for _, am := range allowed {
		if mime == am {
			return true
		}
	}

	return false
}

func ensureBucket(ctx context.Context, bucket string) error {
	exists, err := S3.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}

	if !exists {
		err = S3.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
