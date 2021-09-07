package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/calebhiebert/go-vue-template/db"
	_ "github.com/calebhiebert/go-vue-template/docs"
	"github.com/calebhiebert/go-vue-template/graph"
	"github.com/calebhiebert/go-vue-template/graph/generated"
	"github.com/calebhiebert/go-vue-template/models/modelcrud"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

//go:generate echo "Generating SQL Boilerplate with sqlboiler"
//go:generate sqlboiler --add-global-variants --tag-ignore pw_hash --add-soft-deletes --config ./sqlboiler.toml --templates ./templates --templates $GOPATH/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/templates --templates $GOPATH/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.6.0/templates_test psql
//go:generate echo "Cleaning up SQL models"
//go:generate goimports -w models/
//go:generate echo "Generating GraphQL Code"
//go:generate gqlgen generate
//go:generate echo "Generating goverter code"
//go:generate go run github.com/jmattheis/goverter/cmd/goverter github.com/calebhiebert/go-vue-template/convert
//go:generate echo "Generating swagger documentation"
//go:generate swag init --exclude ui

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
	router.GET("/healthz", c.HealthCheck)
	router.GET("/avatar/:id", c.GenerateAvatar)

	// *******************************
	// * Authenticated Routes        *
	// *******************************
	auth := router.Group("/auth")

	auth.POST("/register", c.RegisterUsernamePassword)
	auth.POST("/loginup", c.AuthenticateUsernamePassword)
	auth.POST("/loginjwt", verifyTokenMiddleware, c.AuthenticateJWT)

	// *******************************
	// * Protected Routes            *
	// *******************************
	protected := router.Group("")
	protected.Use(accessLogMiddleware, verifyTokenMiddleware, mustBeAuthenticatedMiddleware)

	protected.GET("/users/me", c.GetMe)

	// *******************************
	// * Admin Routes                *
	// *******************************
	admin := router.Group("/admin")
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

	config := generated.Config{Resolvers: &graph.Resolver{}}
	config.Directives.HasRole = userHasRoleDirective

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	plgrnd := playground.Handler("GraphQL playground", "/query")

	gql.GET("/graphql", func(c *gin.Context) {
		plgrnd.ServeHTTP(c.Writer, c.Request)
	})

	gql.POST("/query", verifyTokenMiddleware, func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	http.Handle("/query", srv)

	// *******************************
	// * Swagger Setup               *
	// *******************************
	swaggerURL := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", os.Getenv("HOSTED_URL")))
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
