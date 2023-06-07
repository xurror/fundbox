package controllers

import (
	"net/http"
	"strconv"

	"getting-to-go/services"
	"getting-to-go/utils"

	"github.com/gin-gonic/gin"
)

// UserController provides user-related endpoints
type UserController struct {
	userService *services.UserService
}

// NewUserController creates a new UserController instance
func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register registers the UserController routes with the given Gin engine
func (c *UserController) Register(router *gin.RouterGroup) {
	router.POST("/auth", c.authenticate)
	router.POST("/users", c.createUser)
	router.GET("/users/:id", c.getUser)
	router.GET("/api/users", c.getUsers)
}

func (c *UserController) createUser(ctx *gin.Context) {
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

func (c *UserController) authenticate(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}

	user, err := c.userService.Authenticate(req.Email, req.Password)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *UserController) getUsers(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "1"))

	users, err := c.userService.GetUsers(limit, offset)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}

	var res []gin.H
	for _, user := range users {
		res = append(res, gin.H{
			"id":        user.ID,
			"firstName": user.FirstName,
			"lastName":  user.LastName,
			"email":     user.Email,
		})
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetUser(id)

	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":        user.ID,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"email":     user.Email,
		"roles":     user.Roles,
	})
}
