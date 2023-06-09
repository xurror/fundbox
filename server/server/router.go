package server

import (
	controllers "getting-to-go/controller"
	models "getting-to-go/model"
	"getting-to-go/server/middleware"
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
	//r.Use(cors.New(CorsConfig()))

	authMiddleware := middleware.GetAuthMiddleware(models.DB())

	auth := r.Group("/auth")
	new(controllers.AuthController).Register(auth, authMiddleware.LoginHandler)

	r.GET("/graphiql", playgroundHandler())

	// Secured routes
	if !c.DisableAuth {
		r.Use(authMiddleware.MiddlewareFunc())
	}

	r.POST("/auth/refresh_token", authMiddleware.RefreshHandler)
	r.POST("/graphql", graphqlHandler())

	return r
}
