package auth

import (
	"context"
	"getting-to-go/models"
	"getting-to-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware(db *gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		//cookie, err := c.Request.Cookie("auth-cookie")
		jwt := c.Request.Header.Get("Authorization")
		//jwt := c.GetHeader("Authorization")
		//jwt := ""

		// Allow unauthenticated users in
		if jwt == "" || jwt == "null" {
			c.Next()
			return
		}

		userId, err := utils.ValidateToken(jwt)
		if err != nil {
			utils.HandleAppError(c, err)
			return
		}

		// get the user from the database
		user := getUserByID(db, userId)

		// put it in context
		c.Set("userCtxKey", user)

		// and call the next with our new context
		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}

func getUserByID(db *gorm.DB, userId uuid.UUID) *models.User {
	var user models.User
	db.Where("id = ?", userId).First(&user)
	return &user
}
