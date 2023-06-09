package graph

import (
	"context"
	"getting-to-go/model"
	"getting-to-go/util"
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

func HasRolesDirective(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.Role) (res interface{}, err error) {
	ginCtx, err := utils.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if ginCtx == nil {
		log.Panic(ginCtx)
	}
	//claims := jwt.ExtractClaims(ginCtx)
	//user, _ := next.
	//if !middleware.ForContext(ctx).HasRoles(roles) {
	//	return nil, fmt.Errorf("access denied")
	//}
	return next(ctx)
}
