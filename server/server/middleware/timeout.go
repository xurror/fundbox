package middleware

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(5000*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			c.JSON(http.StatusRequestTimeout, gin.H{
				"code":    http.StatusRequestTimeout,
				"message": "Request timeout!",
			})
		}),
	)
}
