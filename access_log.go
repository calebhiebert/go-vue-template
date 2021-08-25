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

func accessLogMiddleware(c *gin.Context) {
	preProcessingSpan, _ := opentracing.StartSpanFromContext(c.Request.Context(), "Access Log Preprocesing")

	accessLog := models.AccessLog{
		ID:        uuid.Must(uuid.NewV4()).String(),
		Path:      c.Request.URL.String(),
		IPAddress: c.ClientIP(),
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

	headerJson, err := json.Marshal(c.Request.Header)
	if err != nil {
		preProcessingSpan.LogKV("event", "error", "message", "failed to marshal incoming headers for access log")
	} else {
		accessLog.RequestHeaders = null.JSONFrom(headerJson)
	}

	c.Request.Body = ioutil.NopCloser(&buf)

	preProcessingSpan.Finish()

	startTime := time.Now().UTC()

	// Create a buffer to store the response body
	var responseBuf bytes.Buffer

	writer := io.MultiWriter(c.Writer, &responseBuf)

	c.Writer = writer

	c.Next()

	durationMS := time.Now().UTC().UnixMilli() - startTime.UnixMilli()

	accessLog.ProcessingDuration = int(durationMS)

	postProcessingAccessLogSpan, postCtx := opentracing.StartSpanFromContext(c.Request.Context(), "Access Log Postprocessing")

	c.JSON()

	err = accessLog.InsertG(postCtx, boil.Infer())
	if err != nil {
		postProcessingAccessLogSpan.LogKV("event", "error", "message", "failed to record access log")
	}
}
