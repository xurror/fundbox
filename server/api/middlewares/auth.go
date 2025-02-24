package middlewares

import (
	"community-funds/config"
	"community-funds/pkg/models"
	"community-funds/pkg/repositories"
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"

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
func AuthMiddleware(userRepo *repositories.UserRepository, cfg *config.Config, log *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		issuerURL, err := url.Parse("https://" + cfg.Auth0.Domain + "/")
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid Auth0 configuration"})
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
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid Auth0 configuration"})
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := v.ValidateToken(c.Context(), tokenString)
		if err != nil {
			log.Debugf("Failed to validate token: %v", err)
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		var userClaims *AuthUserClaims
		claims, claimsOk := token.(*validator.ValidatedClaims)
		if claimsOk {
			userClaims, claimsOk = claims.CustomClaims.(*AuthUserClaims)
		}

		if !claimsOk {
			log.Debug("Failed to extract claims")
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
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
				return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create internal user"})
			}
		}

		// Store user ID in context for handlers
		c.Context().SetUserValue("userID", user.ID)

		return c.Next()
	}
}
