package main

import (
	"embed"
	"fmt"
	gomime "github.com/cubewise-code/go-mime"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path"
)

//go:embed ui/dist/*
var uiFiles embed.FS

func serveUI(c *gin.Context) {
	rPath := c.Request.URL.Path

	// First, check to see if the file being requested exists
	file, err := uiFiles.Open("ui/dist"+rPath)
	if err != nil {
		// The file does not exist, send the index
		file, err = uiFiles.Open("ui/dist/index.html")

		c.Writer.Header().Add("Content-Type", "text/html; charset=utf-8")
		c.Writer.Header().Add("Cache-Control", "no-cache")
		c.Writer.WriteHeader(http.StatusOK)
	} else {
		fileExt := path.Ext(rPath)
		mime := gomime.TypeByExtension(fileExt)

		switch mime {
		case "text/html":
			mime = "text/html; charset=utf-8"
		}

		c.Writer.Header().Add("Content-Type", mime)
		c.Writer.Header().Add("Cache-Control", "no-cache") // TODO proper cache rules
		c.Writer.WriteHeader(http.StatusOK)
	}

	_, err = io.Copy(c.Writer, file)
	if err != nil {
		fmt.Println("Failed to send the response file to the client", err.Error())
	}
}
