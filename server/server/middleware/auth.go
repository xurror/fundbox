package middleware

import (
	"getting-to-go/model"
	"getting-to-go/util"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var identityKey = "id"

func GetAuthMiddleware(db *gorm.DB) *jwt.GinJWTMiddleware {
	m, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := uuid.MustParse(claims[identityKey].(string))
			user := &models.User{}
			db.First(&user, "id = ?", id)
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login Login
			if err := c.ShouldBindJSON(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user := &models.User{}
			result := db.First(&user, "email = ?", login.Email)
			if result.Error != nil {
				return nil, utils.NewError(http.StatusUnauthorized, jwt.ErrFailedAuthentication.Error())
			}

			// If the password does not match, return an error
			if !utils.CheckPasswordHash(user.Password, login.Password) {
				return nil, utils.NewError(http.StatusUnauthorized, "Invalid email or password")
			}
			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.HasRoles([]models.Role{models.Initiator, models.Contributor}) {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := m.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return m
}
