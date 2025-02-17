package handlers

import (
	"community-funds/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users := h.Service.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"users": users})
}
