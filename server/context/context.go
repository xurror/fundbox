package context

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
)

var contextKey = "AppContext"

func AcquireAppContext(c context.Context) echo.Context {
	ctx := c.Value(contextKey)
	if ctx == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		panic(err)
	}

	ec, ok := ctx.(echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		panic(err)
	}
	return ec
}

func RegisterAppContext() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.WithValue(c.Request().Context(), contextKey, c)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
