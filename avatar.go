package main

import (
	"fmt"
	"image/jpeg"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/o1egl/govatar"
)

func (*Controller) GenerateAvatar(c *gin.Context) {
	id := c.Param("id")

	img, err := govatar.GenerateForUsername(govatar.MALE, id)
	if err != nil {
		APIErrorFromErr(err).Respond(c)
		return
	}

	c.Writer.Header().Set("Content-Type", "image/jpeg")
	c.Writer.Header().Set("Cache-Control", "public, max-age=604800, immutable")
	c.Writer.WriteHeader(http.StatusOK)

	err = jpeg.Encode(c.Writer, img, &jpeg.Options{Quality: 50})
	if err != nil {
		fmt.Println("ERROR", err)
	}
}
