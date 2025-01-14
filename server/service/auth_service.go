package service

import (
	"context"
	"encoding/json"
	"fmt"
	"getting-to-go/config"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Scope string `json:"scope"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

type AuthService struct {
	logger      *logrus.Logger
	config      *config.AppConfig
	userService UserService
}

func NewAuthService(logger *logrus.Logger, config *config.AppConfig, userService UserService) *AuthService {
	return &AuthService{
		logger:      logger,
		config:      config,
		userService: userService,
	}
}

func (s *AuthService) unmarshalJson(bytes []byte) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	return data
}

func (s *AuthService) unmarshalResponse(res *http.Response) map[string]interface{} {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.Errorf("Failed to read response body: %v", err)
		return nil
	}
	return s.unmarshalJson(body)
}

func (s *AuthService) executeRequest(req *http.Request) map[string]interface{} {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Errorf("Failed to get response: %v", err)
		return nil
	}
	defer res.Body.Close()
	return s.unmarshalResponse(res)
}

func (s *AuthService) getManagementRequestBody() string {
	data := map[string]string{
		"client_id":     s.config.Auth0.Management.ClientId,
		"client_secret": s.config.Auth0.Management.ClientSecret,
		"audience":      "https://" + s.config.Auth0.Management.Domain + "/api/v2/",
		"grant_type":    "client_credentials",
	}
	body, err := json.Marshal(data)
	if err != nil {
		s.logger.Fatalf("Failed to marshal request body: %v", err)
	}
	return string(body)
}

func (s *AuthService) getAuth0ManagementToken(auth0Id string) {
	req, _ := http.NewRequest(
		"POST",
		"https://"+s.config.Auth0.Management.Domain+"/oauth/token",
		strings.NewReader(s.getManagementRequestBody()),
	)
	req.Header.Add("content-type", "application/json")

	data := s.executeRequest(req)
	fmt.Println(data)

	req, _ = http.NewRequest(
		"GET",
		"https://"+s.config.Auth0.Management.Domain+"/api/v2/users/"+auth0Id,
		nil,
	)
	req.Header.Add("Authorization", "Bearer "+data["access_token"].(string))

	data = s.executeRequest(req)
	fmt.Println(data)
}

func (s *AuthService) Authorize(ctx echo.Context, tokenString string) (interface{}, error) {

	issuerURL, err := url.Parse("https://" + s.config.Auth0.Domain + "/")
	if err != nil {
		s.logger.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{s.config.Auth0.Audience},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		s.logger.Fatalf("Failed to set up the jwt validator")
	}

	token, err := jwtValidator.ValidateToken(ctx.Request().Context(), tokenString)
	if err != nil {
		s.logger.Debug(err.Error())
		return nil, echo.ErrUnauthorized
	}

	fmt.Print(token)
	if claims, ok := token.(*validator.ValidatedClaims); ok {
		auth0Id := claims.RegisteredClaims.Subject
		_, err := s.userService.GetUserByAuth0Id(ctx.Request().Context(), auth0Id)
		if err != nil {
			go func() {
				_, err := s.userService.CreateUser(context.Background(), auth0Id)
				if err != nil {
					s.logger.Debug(err.Error())
				}
			}()
		}

		return claims.RegisteredClaims, nil
	} else {
		s.logger.Debug("Invalid token claims")
		return nil, echo.ErrUnauthorized
	}
}
