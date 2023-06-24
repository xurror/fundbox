package graph

import (
	"context"
	"getting-to-go/model"
	"github.com/99designs/gqlgen/graphql"
	"log"
)

//func HasRolesDirective(ctx context.Context, obj interface{}, next graphql.Resolver, roles []string) (res interface{}, err error) {
//	// Get the current user from the context
//	user := ctx.Value("user")
//	if user == nil {
//		return nil, gqlerror.Errorf("Access denied!")
//	}
//
//	// Get the user roles
//	userRoles := user.(*models.User).Roles
//
//	// Check if the user has the required roles
//	if utils.ContainsAny(userRoles, roles) {
//		return next(ctx)
//	}
//
//	return nil, gqlerror.Errorf("Access denied!")
//}

func HasRolesDirective(ctx context.Context, obj interface{}, next graphql.Resolver, roles []model.Role) (res interface{}, err error) {
	user := ctx.Value("user").(*model.User)
	log.Panic(user)
	//claims := jwt.ExtractClaims(ginCtx)
	//user, _ := next.
	//if !middleware.ForContext(ctx).HasRoles(roles) {
	//	return nil, fmt.Errorf("access denied")
	//}
	return next(ctx)
}
