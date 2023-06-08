package main

import (
	"context"
	"fmt"
	"getting-to-go/auth"
	"getting-to-go/config"
	"getting-to-go/controllers"
	"getting-to-go/graph/generated"
	"getting-to-go/graph/resolvers"
	"getting-to-go/models"
	"getting-to-go/services"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := generated.Config{Resolvers: &resolvers.Resolver{}}
	c.Directives.HasRoles = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.Role) (res interface{}, err error) {
		if !auth.ForContext(ctx).HasRoles(roles) {
			return nil, fmt.Errorf("access denied")
		}
		return next(ctx)
	}
	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	// Panic recovery or HandlerFunc
	h.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		// notify bug tracker...
		return gqlerror.Errorf("Internal server error!")
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	fmt.Println("Hello, world!")

	appConfig, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = models.Connect(appConfig.Database.Host, appConfig.Database.Port, appConfig.Database.User, appConfig.Database.Password, appConfig.Database.Name)
	if err != nil {
		fmt.Println(err)
	}

	models.RunMigrations()

	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	fundService := services.NewFundService()
	fundController := controllers.NewFundController(fundService)

	contributionService := services.NewContributionService()
	contributionController := controllers.NewContributionController(contributionService)

	router := gin.Default()

	api := router.Group("/api")
	apiV1 := api.Group("/v1")

	userController.Register(apiV1)
	fundController.Register(apiV1)
	contributionController.Register(apiV1)

	router.Use(auth.Middleware(models.DB()))

	router.POST("/graphql/query", graphqlHandler())
	router.GET("/graphql", playgroundHandler())

	err = router.Run(fmt.Sprintf(":%s", appConfig.Server.Port))
	if err != nil {
		return
	}
}
