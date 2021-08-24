package main

import (
	"fmt"
	"os"
	"time"

	"github.com/calebhiebert/go-vue-template/db"
	_ "github.com/calebhiebert/go-vue-template/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// @title Go Vue Template
// @version 0.1

// @license.name Unknown

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

	boil.SetDB(dbConn)

	ginEngine := gin.Default()

	ginEngine.Use(ginhttp.Middleware(tracer))

	// TODO update with proper cors config
	ginEngine.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	c := NewController()

	ginEngine.GET("/healthz", c.HealthCheck)
	ginEngine.GET("/test", c.Test)

	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json")))

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
