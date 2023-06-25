package server

import (
	"context"
	"getting-to-go/graph"
	"getting-to-go/graph/generated"
	"getting-to-go/graph/resolver"
	"getting-to-go/service"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
)

type GraphQlHandler struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewGraphQlHandler(
	userService *service.UserService,
	authService *service.AuthService,
) *GraphQlHandler {
	return &GraphQlHandler{
		userService: userService,
		authService: authService,
	}
}

func (*GraphQlHandler) GraphiQlHandler(name, pattern string) echo.HandlerFunc {
	h := playground.Handler(name, pattern)
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func (g *GraphQlHandler) QueryHandler() echo.HandlerFunc {
	c := generated.Config{Resolvers: &resolver.Resolver{
		UserService: g.userService,
		AuthService: g.authService,
	}}

	c.Directives.HasRoles = graph.HasRolesDirective

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	h.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})
	h.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		// notify bug tracker...
		log.Print(err)
		switch v := err.(type) {
		case string:
			return gqlerror.Errorf(v)
		case error:
			return gqlerror.Errorf(v.Error())
		default:
			return gqlerror.Errorf("internal server error")
		}
	})

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
