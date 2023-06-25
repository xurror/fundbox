package service

import (
	"getting-to-go/config"
	"getting-to-go/model"
	_type "getting-to-go/type"
	"getting-to-go/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AuthService struct {
	log         *logrus.Logger
	config      *config.AppConfig
	userService *UserService
}

func NewAuthService(log *logrus.Logger, config *config.AppConfig, userService *UserService) *AuthService {
	return &AuthService{
		log:         log,
		config:      config,
		userService: userService,
	}
}

func (s *AuthService) Authenticate(email, password string) (*model.User, error) {
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		log.Debug(err.Error())
		return nil, err
	}

	if user == nil {
		return nil, _type.NewAppError(http.StatusUnauthorized, "invalid email or password")
	}

	if !util.CheckPasswordHash(user.Password, password) {
		return nil, _type.NewAppError(http.StatusUnauthorized, "invalid email or password")
	}

	return user, nil
}

func (s *AuthService) Authorize(c echo.Context, tokenString string) (interface{}, error) {
	var options []jwt.ParseOption
	options = append(options, jwt.WithToken(jwt.New()))
	options = append(options, jwt.WithVerify(true))
	options = append(options, jwt.WithKey(jwa.HS256, []byte(s.config.Jwt.SigningKey)))
	options = append(options, jwt.WithValidate(true))

	token, err := jwt.ParseString(tokenString, options...)
	if err != nil {
		log.Debug(err.Error())
		return nil, echo.ErrUnauthorized
	}

	user, err := s.userService.GetUserByEmail(token.Subject())
	if err != nil {
		log.Debug(err.Error())
		return nil, echo.ErrUnauthorized
	}
	return user, nil
}

func (s *AuthService) SignUp(firstName, lastName, email, password string, roles []model.Role) (*model.User, error) {
	_, err := s.userService.GetUserByEmail(email)
	if err == nil {
		return nil, _type.NewAppError(http.StatusConflict, "email already exists")
	}

	return s.userService.CreateUser(firstName, lastName, email, password, roles)
}
