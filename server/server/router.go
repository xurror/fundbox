package server

import (
	"getting-to-go/controller"
	"getting-to-go/model"
	"getting-to-go/server/middleware"
	"getting-to-go/service"

	"github.com/gin-contrib/location"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func NewRouter(c *RouterConfig) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.GinContextToContextMiddleware())
	r.Use(requestid.New())
	r.Use(location.Default())
	r.Use(middleware.TimeoutMiddleware())
	//r.Use(secure.New(SecureConfig()))
	// r.Use(cors.New(CorsConfig()))

	authMiddleware := middleware.GetAuthMiddleware(model.DB())
	userService := service.NewUserService()
	authService := service.NewAuthService(userService)
	authController := controller.NewAuthController(authService)

	auth := r.Group("/auth")
	authController.Register(auth, authMiddleware.LoginHandler)

	r.GET("/graphiql", playgroundHandler())

	// Secured routes
	if !c.DisableAuth {
		r.Use(authMiddleware.MiddlewareFunc())
	}

	r.POST("/auth/refresh_token", authMiddleware.RefreshHandler)
	r.POST("/graphql", graphqlHandler())

	return r
}
