package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"time"

	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type responseBodyWrapper struct {
	gin.ResponseWriter
	body       *bytes.Buffer
}

func (r responseBodyWrapper) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func accessLogMiddleware(c *gin.Context) {
	preProcessingSpan, _ := opentracing.StartSpanFromContext(c.Request.Context(), "Access Log Preprocesing")

	accessLog := models.AccessLog{
		ID:            uuid.Must(uuid.NewV4()).String(),
		Path:          c.Request.URL.String(),
		RequestMethod: c.Request.Method,
		IPAddress:     c.ClientIP(),
	}

	// Read the incoming request body
	var buf bytes.Buffer

	tee := io.TeeReader(c.Request.Body, &buf)

	body, err := ioutil.ReadAll(tee)
	if err != nil {
		preProcessingSpan.LogKV("event", "error", "message", "failed to read incoming request body for access log")
	} else {
		accessLog.RequestBody = null.StringFrom(string(body))
	}

	clonedRequestHeaders := c.Request.Header.Clone()

	// Remove the authorization header from logging
	clonedRequestHeaders.Del("Authorization")

	headerJson, err := json.Marshal(clonedRequestHeaders)
	if err != nil {
		preProcessingSpan.LogKV("event", "error", "message", "failed to marshal incoming headers for access log")
	} else {
		accessLog.RequestHeaders = null.JSONFrom(headerJson)
	}

	c.Request.Body = ioutil.NopCloser(&buf)

	preProcessingSpan.Finish()

	w := &responseBodyWrapper{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
	c.Writer = w

	startTime := time.Now().UTC()

	c.Next()

	durationMS := time.Now().UTC().UnixMilli() - startTime.UnixMilli()

	accessLog.ProcessingDuration = int(durationMS)

	postProcessingAccessLogSpan, postCtx := opentracing.StartSpanFromContext(c.Request.Context(), "Access Log Postprocessing")
	defer postProcessingAccessLogSpan.Finish()

	err = accessLog.ResponseBody.UnmarshalJSON(w.body.Bytes())
	if err != nil {
		postProcessingAccessLogSpan.LogKV("event", "error", "message", "failed to unmarshal response body for access log")

		accessLog.ResponseBody.Marshal(map[string]interface{}{
			"error": "response body not json",
		})
	}

	err = accessLog.ResponseHeaders.Marshal(w.ResponseWriter.Header())
	if err != nil {
		postProcessingAccessLogSpan.LogKV("event", "error", "message", "failed to marshal response headers for access log")
	}

	user := extractVerifiedUser(c)

	if user != nil {
		accessLog.UserID = null.StringFrom(user.ID)
	}

	accessLog.ResponseCode = w.ResponseWriter.Status()

	err = accessLog.InsertG(postCtx, boil.Infer())
	if err != nil {
		postProcessingAccessLogSpan.LogKV("event", "error", "message", "failed to record access log")
	}
}
