package controller

import (
	"getting-to-go/config"
	"getting-to-go/model"
	"getting-to-go/service"
	_type "getting-to-go/type"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type AuthController struct {
	log         *logrus.Logger
	config      *config.AppConfig
	authService *service.AuthService
}

func NewAuthController(log *logrus.Logger, config *config.AppConfig, authService *service.AuthService) *AuthController {
	return &AuthController{
		log:         log,
		config:      config,
		authService: authService,
	}
}

func (a *AuthController) Register(g *echo.Group) {
	g.POST("/login", a.login)
	g.POST("/signup", a.signUp)
}

func (a *AuthController) login(c echo.Context) error {
	req := &LoginRequest{}
	if err := c.Bind(req); err != nil {
		log.Debug(err.Error())
		return echo.ErrBadRequest
	}

	user, err := a.authService.Authenticate(req.Email, req.Password)
	if err != nil {
		log.Debug(err.Error())
		return echo.ErrUnauthorized
	}

	now := time.Now()

	token, err := jwt.NewBuilder().
		Issuer("fundbox").
		IssuedAt(now).
		Expiration(now.Add(a.config.Jwt.Expiration)).
		Subject(user.Email).
		JwtID(user.ID.String()).
		Build()
	if err != nil {
		log.Debug(err.Error())
		return echo.ErrInternalServerError
	}

	tokenString, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, []byte(a.config.Jwt.SigningKey)))
	if err != nil {
		log.Debug(err.Error())
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, &JwtResponse{
		Code:   http.StatusOK,
		Expire: token.Expiration(),
		Token:  string(tokenString[:]),
	})
}

func (a *AuthController) signUp(c echo.Context) error {
	req := &SignUpRequest{}
	if err := c.Bind(req); err != nil {
		log.Debug(err.Error())
		return echo.ErrBadRequest
	}

	user, err := a.authService.SignUp(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
		req.Roles,
	)
	if err != nil {
		log.Debug(err.Error())
		switch e := err.(type) {
		case *_type.AppError:
			return echo.NewHTTPError(e.Code, e.Message)
		default:
			return echo.ErrInternalServerError
		}
	}

	return c.JSON(http.StatusCreated, &user)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	FirstName string       `json:"first_name" binding:"required"`
	LastName  string       `json:"last_name" binding:"required"`
	Email     string       `json:"email" binding:"required,email"`
	Password  string       `json:"password" binding:"required,min=6"`
	Roles     []model.Role `json:"roles" binding:"required"`
}

type JwtResponse struct {
	Code   int       `json:"code"`
	Expire time.Time `json:"expire"`
	Token  string    `json:"token"`
}
