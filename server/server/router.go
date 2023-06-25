package server

import (
	appContext "getting-to-go/context"
	"getting-to-go/controller"
	"getting-to-go/service"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func NewRouter(
	log *logrus.Logger,
	graphQl *GraphQlHandler,
	authService *service.AuthService,
	authController *controller.AuthController,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.RequestLoggerWithConfig(RequestLoggerConfig(log)))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(CorsConfig()))
	e.Use(middleware.TimeoutWithConfig(TimeoutConfig()))
	e.Use(middleware.BodyLimit("2M"))
	e.Use(appContext.RegisterAppContext())
	//e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	//}))

	authGroup := e.Group("/auth")
	authController.Register(authGroup)

	e.GET("/graphiql", graphQl.GraphiQlHandler("Fund Box", "/gql/query"))
	graphql := e.Group("/gql", echojwt.WithConfig(JwtConfig(authService)))
	graphql.POST("/query", graphQl.QueryHandler())

	return e
}
