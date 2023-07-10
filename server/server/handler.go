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
	userService         *service.UserService
	authService         *service.AuthService
	fundService         *service.FundService
	contributionService *service.ContributionService
}

func NewGraphQlHandler(
	userService *service.UserService,
	authService *service.AuthService,
	fundService *service.FundService,
	contributionService *service.ContributionService,
) *GraphQlHandler {
	return &GraphQlHandler{
		userService:         userService,
		authService:         authService,
		fundService:         fundService,
		contributionService: contributionService,
	}
}

func (*GraphQlHandler) GraphiQlHandler(name, pattern string) echo.HandlerFunc {
	srv := playground.Handler(name, pattern)
	return echo.WrapHandler(srv)
}

func (g *GraphQlHandler) QueryHandler() echo.HandlerFunc {
	c := generated.Config{Resolvers: &resolver.Resolver{
		UserService:         g.userService,
		AuthService:         g.authService,
		FundService:         g.fundService,
		ContributionService: g.contributionService,
	}}

	c.Directives.HasRoles = graph.HasRolesDirective

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
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

	return echo.WrapHandler(srv)
}
