package server

import (
	"context"
	"getting-to-go/graph"
	"getting-to-go/graph/generated"
	"getting-to-go/graph/resolver"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func graphqlHandler() gin.HandlerFunc {
	c := generated.Config{Resolvers: &resolver.Resolver{}}

	c.Directives.HasRoles = graph.HasRolesDirective

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))
	h.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})
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
	h := playground.Handler("GraphQL Playground", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
