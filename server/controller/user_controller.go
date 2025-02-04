package controller

import (
	"getting-to-go/config"
	"getting-to-go/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	logger      *logrus.Logger
	config      *config.AppConfig
	userService model.UserService
}

func NewUserController(logger *logrus.Logger, config *config.AppConfig, userService model.UserService) *UserController {
	return &UserController{
		logger:      logger,
		config:      config,
		userService: userService,
	}
}

func (a *UserController) RegisterRoutes(g *echo.Group) {
	g.POST("/getUsers", a.getUsers)
}

func (a *UserController) getUsers(c echo.Context) error {
	req := &LoginRequest{}
	if err := c.Bind(req); err != nil {
		log.Debug(err.Error())
		return echo.ErrBadRequest
	}

	user, _, err := a.userService.GetUsers(c.Request().Context(), 10, nil)
	if err != nil {
		log.Debug(err.Error())
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, &user)
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
