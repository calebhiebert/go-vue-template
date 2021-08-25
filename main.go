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

	// Load the .env file (if present)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env file. Most of time this error can be ignored")
	}

	// Create a opentracing tracer setup to send data to jaeger
	tracer, closer, err := SetupTracing()
	if err != nil {
		panic(err)
	}

	// Defer closing of the tracer
	defer closer.Close()

	// Create a connection to the postgres database
	dbConn, err := db.SetupDatabase()
	if err != nil {
		panic(err)
	}

	// Setup the global db connection for sqlboiler
	boil.SetDB(dbConn)

	// Create the gin router
	router := gin.Default()

	// Use the tracing middleware so all incoming requests get traced
	router.Use(ginhttp.Middleware(tracer))

	// Setup CORS
	// TODO update with proper cors config
	router.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	// Create a new controller (this is where the api handlers live)
	c := NewController()

	router.GET("/healthz", c.HealthCheck)
	router.GET("/test", c.Test)

	swaggerURL := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", os.Getenv("HOSTED_URL")))

	// Setup the route for swagger ui serving
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// Default port 8080, but check if an env port should override it
	port := "8080"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Run the app
	err = router.Run(fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		fmt.Println("ERROR starting server", err.Error())
	}
}
