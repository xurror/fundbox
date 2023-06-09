package controller

import (
	"getting-to-go/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController provides user-related endpoints
type AuthController struct {
	authService *service.AuthService
}

// NewAuthController creates a new AuthController instance
func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register registers the AuthController routes with the given Gin engine
func (c *AuthController) Register(router *gin.RouterGroup, loginHandler func(c *gin.Context)) {
	router.POST("/signup", c.signUp)
	router.POST("/login", loginHandler)
}

func (c *AuthController) signUp(ctx *gin.Context) {
	var req struct {
		FirstName string   `json:"first_name" binding:"required"`
		LastName  string   `json:"last_name" binding:"required"`
		Email     string   `json:"email" binding:"required,email"`
		Password  string   `json:"password" binding:"required,min=6"`
		Roles     []string `json:"roles" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Panic(err)
	}

	user, err := c.authService.SignUp(
		ctx,
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
		req.Roles,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		log.Panic(err)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"roles":     user.Roles,
	})
}
