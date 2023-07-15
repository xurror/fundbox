package context

import (
	"context"
	"fmt"
	"getting-to-go/model"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var contextKey = "AppContext"

func AcquireAppContext(ctx context.Context) echo.Context {
	c := ctx.Value(contextKey)
	if c == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		panic(err)
	}

	eCtx, ok := c.(echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		panic(err)
	}
	return eCtx
}

func RegisterAppContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			ctx := context.WithValue(eCtx.Request().Context(), contextKey, eCtx)
			eCtx.SetRequest(eCtx.Request().WithContext(ctx))
			return next(eCtx)
		}
	}
}

func CurrentUser(ctx context.Context) *model.User {
	eCtx := AcquireAppContext(ctx)
	user := eCtx.Get("user")
	if user == nil {
		err := fmt.Errorf("could not retrieve user from context")
		log.Panic(err)
	}
	return (user).(*model.User)
}
