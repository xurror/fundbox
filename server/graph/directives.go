package graph

import (
	"context"
	appContext "getting-to-go/context"
	"getting-to-go/model"
	"github.com/99designs/gqlgen/graphql"
	log "github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func HasRolesDirective(ctx context.Context, obj interface{}, next graphql.Resolver, roles []model.Role) (res interface{}, err error) {
	user := appContext.CurrentUser(ctx)
	if user.HasRoles(roles) {
		return next(ctx)
	}

	log.Debug("Insufficient Permissions")
	return nil, gqlerror.Errorf("Access denied!")
}
