package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetCurrentUserID(c *fiber.Ctx) *uuid.UUID {
	id, ok := c.Context().UserValue("userID").(uuid.UUID)
	if !ok {
		return nil
	}
	return &id
}
