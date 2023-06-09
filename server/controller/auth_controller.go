package controllers

import (
	"getting-to-go/service"
	"getting-to-go/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController provides user-related endpoints
type AuthController struct {
	userService *services.UserService
}

// NewAuthController creates a new AuthController instance
func NewAuthController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register registers the UserController routes with the given Gin engine
func (c *AuthController) Register(router *gin.RouterGroup, loginHandler func(c *gin.Context)) {
	router.POST("/login", loginHandler)
	router.POST("/signup", c.createUser)
}

func (c *AuthController) createUser(ctx *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}

	user, err := c.userService.CreateUser(req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"roles":     user.Roles,
	})
}
