package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/calebhiebert/go-vue-template/db"
	"github.com/calebhiebert/go-vue-template/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/opentracing-contrib/go-gin/ginhttp"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env file. Most of time this error can be ignored")
	}

	tracer, closer, err := SetupTracing()
	if err != nil {
		panic(err)
	}

	initSpan := tracer.StartSpan("Initialization")

	defer closer.Close()

	dbConn, err := db.SetupDatabase()
	if err != nil {
		panic(err)
	}

	ginEngine := gin.Default()

	ginEngine.Use(ginhttp.Middleware(tracer))

	ginEngine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	ginEngine.GET("/test", func(c *gin.Context) {
		test, err := models.Users().All(c.Request.Context(), dbConn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		test2, err := models.TestTables().One(c.Request.Context(), dbConn)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"test": test, "test2": test2})
	})

	port := "8080"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	initSpan.Finish()

	err = ginEngine.Run(fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		fmt.Println("ERROR starting server", err.Error())
	}
}