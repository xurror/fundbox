package handlers

import (
	"community-funds/api/dto"
	"community-funds/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// GetCurrentUser retrieves current logged in user
// @Summary Get current authenticated user
// @Description Returns the current authenticated user id
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} dto.UserDTO
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /users/me [get]
func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
	currentUserID := utils.GetCurrentUserID(c)
	if currentUserID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.UserDTO{ID: *currentUserID})
}
