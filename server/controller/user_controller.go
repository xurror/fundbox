package controller

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"

	"getting-to-go/service"
	"getting-to-go/util"

	"github.com/gin-gonic/gin"
)

// UserController provides user-related endpoints
type UserController struct {
	userService *service.UserService
}

// NewUserController creates a new UserController instance
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register registers the UserController routes with the given Gin engine
func (c *UserController) Register(router *gin.RouterGroup) {
	router.POST("/users", c.createUser)
	router.GET("/users/:id", c.getUser)
	router.GET("/api/users", c.getUsers)
}

func (c *UserController) createUser(ctx *gin.Context) {
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

	user, err := c.userService.CreateUser(req.FirstName, req.LastName, req.Email, req.Password, req.Roles)
	if err != nil {
		util.HandleAppError(ctx, err)
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

func (c *UserController) getUsers(ctx *gin.Context) {
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "1"))

	users, err := c.userService.GetUsers(limit, offset)
	if err != nil {
		util.HandleAppError(ctx, err)
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
	user, err := c.userService.GetUser(uuid.MustParse(id))

	if err != nil {
		util.HandleAppError(ctx, err)
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
