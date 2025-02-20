package middlewares

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"community-funds/internal/config"
	"community-funds/internal/models"
	"community-funds/internal/repositories"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AuthClaims contains custom data we want from the token.
type AuthUserClaims struct {
	Scope string `json:"scope"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c AuthUserClaims) Validate(ctx context.Context) error {
	return nil
}

// AuthMiddleware enforces authentication and maps Auth0 users to internal users
func AuthMiddleware(userRepo *repositories.UserRepository, cfg *config.Config, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		issuerURL, err := url.Parse("https://" + cfg.Auth0.Domain + "/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Auth0 configuration"})
			c.Abort()
			return
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		v, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{cfg.Auth0.Audience},
			validator.WithCustomClaims(
				func() validator.CustomClaims {
					return &AuthUserClaims{}
				},
			),
			validator.WithAllowedClockSkew(time.Minute),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Auth0 configuration"})
			c.Abort()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := v.ValidateToken(c.Request.Context(), tokenString)
		if err != nil {
			log.Debugf("Failed to validate token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		var userClaims *AuthUserClaims
		claims, claimsOk := token.(*validator.ValidatedClaims)
		if claimsOk {
			userClaims, claimsOk = claims.CustomClaims.(*AuthUserClaims)
		}

		if !claimsOk {
			log.Debug("Failed to extract claims")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		auth0ID := claims.RegisteredClaims.Subject // Extract Auth0 User ID (`sub` claim)

		// Find or Create Internal User
		user, err := userRepo.GetUserByAuth0ID(auth0ID)
		if err != nil {

			user = &models.User{
				Auth0ID: auth0ID,
				Name:    userClaims.Name,
				Email:   userClaims.Email,
			}

			if err := userRepo.CreateUser(user); err != nil {
				log.Debugf("Failed to create internal user: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create internal user"})
				c.Abort()
				return
			}
		}

		// Store user ID in context for handlers
		c.Set("user_id", user.ID)

		c.Next()
	}
}
