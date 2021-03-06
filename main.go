package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/calebhiebert/go-vue-template/graph/resolve"
	"github.com/calebhiebert/go-vue-template/jobs"
	"github.com/calebhiebert/go-vue-template/s3"
	lggr "github.com/datomar-labs-inc/FCT_Helpers_Go/logger"
	fct_tracing "github.com/datomar-labs-inc/FCT_Helpers_Go/tracing"
	limits "github.com/gin-contrib/size"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/calebhiebert/go-vue-template/db"
	_ "github.com/calebhiebert/go-vue-template/docs"
	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/models/modelcrud"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate echo "Generating SQL Boilerplate with sqlboiler"
//go:generate sqlboiler --add-global-variants --tag-ignore pw_hash --add-soft-deletes --config ./sqlboiler.toml --templates ./templates --templates $GOPATH/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/templates --templates $GOPATH/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/templates_test psql
//go:generate echo "Cleaning up SQL models"
//go:generate goimports -w models/
//go:generate echo "Generating GraphQL Code"
//go:generate gqlgen generate --verbose
//go:generate echo "Generating swagger documentation"
//go:generate swag init --exclude ui

// @title Go Vue Template
// @version 0.1

// @license.name Unknown

var ServiceName = "GoVueTemplate"

func main() {

	// Load the .env file (if present)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load env file. Most of time this error can be ignored")
	}

	// Create a tracer
	traceCloser, err := fct_tracing.SetupTracing(ServiceName, trace.TraceIDRatioBased(getTraceRatio()))
	if err != nil {
		lggr.Get("setup-tracing").Fatal("failed to set up tracer", zap.Error(err))
	}

	defer traceCloser()

	// Create a connection to the postgres database
	dbConn, err := db.SetupDatabase()
	if err != nil {
		panic(err)
	}

	// Setup the global db connection for sqlboiler
	boil.SetDB(dbConn)

	// Setup the connection to s3
	err = s3.InitS3()
	if err != nil {
		panic(err)
	}

	err = jobs.InitJobs()
	if err != nil {
		panic(err)
	}

	// Create the gin router
	router := gin.Default()

	// Use the tracing middleware so all incoming requests get traced
	router.Use(otelgin.Middleware(ServiceName, otelgin.WithTracerProvider(otel.GetTracerProvider())))

	// Setup CORS
	// TODO update with proper cors config
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	ConfigAdminCrudModels()

	// Create a new controller (this is where the api handlers live)
	c := NewController()

	// *******************************
	// * Unprotected Routes          *
	// *******************************
	router.GET("/api/healthz", c.HealthCheck)
	router.GET("/api/avatar/:id", c.GenerateAvatar)

	router.GET("/api/image/:id", s3.GetImage)
	router.POST("/api/image", limits.RequestSizeLimiter(s3.MaxImageSize), s3.UploadImage)

	// *******************************
	// * Authenticated Routes        *
	// *******************************
	auth := router.Group("/auth")

	auth.POST("/register", c.RegisterUsernamePassword)
	auth.POST("/loginup", c.AuthenticateUsernamePassword)
	auth.POST("/loginjwt", c.AuthenticateJWT)

	// *******************************
	// * Protected Routes            *
	// *******************************
	protected := router.Group("")
	protected.Use(accessLogMiddleware, verifyTokenMiddleware, mustBeAuthenticatedMiddleware)

	protected.GET("/api/users/me", c.GetMe)

	// *******************************
	// * Admin Routes                *
	// *******************************
	admin := router.Group("/api/admin")
	admin.Use(accessLogMiddleware, verifyTokenMiddleware, mustBeAuthenticatedMiddleware, userHasRoleMiddleware("admin"))

	admin.GET("/dashStats", GetAdminDashboardStats)

	crud := router.Group("/crud")

	crud.Use(accessLogMiddleware, verifyTokenMiddleware, mustBeAuthenticatedMiddleware, userHasRoleMiddleware("admin"))
	modelcrud.RegisterGeneratedCrud(crud)

	// *******************************
	// * Graphql Setup               *
	// *******************************
	gql := router.Group("")

	gql.Use(accessLogMiddleware)

	config := generated.Config{Resolvers: &resolve.Resolver{}}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	plgrnd := playground.Handler("GraphQL playground", "/gql")

	gql.GET("/graphql", func(c *gin.Context) {
		plgrnd.ServeHTTP(c.Writer, c.Request)
	})

	gql.POST("/gql/mutate", verifyTokenMiddleware, func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	http.Handle("/gql/mutate", srv)

	// *******************************
	// * Swagger Setup               *
	// *******************************
	swaggerURL := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", os.Getenv("HOSTED_URL")))
	router.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))

	// *******************************
	// * UI Embed Setup              *
	// *******************************
	router.Use(serveUI)

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

func getTraceRatio() float64 {
	tr := os.Getenv("TRACE_RATIO")
	if tr == "" {
		return 1
	}

	float, err := strconv.ParseFloat(tr, 64)
	if err != nil {
		log.Fatal("Invalid TRACE_RATIO env value set")
	}

	return float
}
