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
)

// AuthClaims contains custom data we want from the token.
type AuthClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c AuthClaims) Validate(ctx context.Context) error {
	return nil
}

// AuthMiddleware enforces authentication and maps Auth0 users to internal users
func AuthMiddleware(userRepo *repositories.UserRepository, cfg *config.Config) gin.HandlerFunc {
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
					return &AuthClaims{}
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.(*validator.ValidatedClaims); ok {
			auth0ID := claims.RegisteredClaims.Subject // Extract Auth0 User ID (`sub` claim)

			// Find or Create Internal User
			user, err := userRepo.GetUserByAuth0ID(auth0ID)
			if err != nil {
				user = &models.User{
					Auth0ID: auth0ID,
					Name:    "claims[\"name\"].(string)",
					Email:   "claims[\"email\"].(string)",
					Role:    "contributor", // Default role
				}
				if err := userRepo.CreateUser(user); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create internal user"})
					c.Abort()
					return
				}
			}

			// Store user ID in context for handlers
			c.Set("user_id", user.ID)
			c.Set("user_role", user.Role)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

	}
}
