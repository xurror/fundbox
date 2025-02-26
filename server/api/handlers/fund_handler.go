package handlers

import (
	"community-funds/api/dto"
	"community-funds/pkg/services"
	"community-funds/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// FundHandler handles fund-related routes
type FundHandler struct {
	service *services.FundService
}

func NewFundHandler(s *services.FundService) *FundHandler {
	return &FundHandler{
		service: s,
	}
}

// CreateFund handles fund creation
// @Summary Create a fund
// @Description Fund managers can create a fund with a target amount
// @Tags funds
// @Accept json
// @Produce json
// @Success 201 {object} dto.FundDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /funds [post]
func (h *FundHandler) CreateFund(c *fiber.Ctx) error {
	var req struct {
		Name         string  `json:"name" binding:"required"`
		TargetAmount float64 `json:"targetAmount" binding:"required"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	managerID := utils.GetCurrentUserID(c)
	if managerID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	fund, err := h.service.CreateFund(req.Name, *managerID, req.TargetAmount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create fund"})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.MapFundToDTO(*fund))
}

// GetFunds retrieves all funds managed by the authenticated user
// @Summary Get all funds managed by the authenticated user
// @Description Returns a list of all funds where the authenticated user is the manager
// @Tags funds
// @Accept json
// @Produce json
// @Query contributorId path string false "Contributor ID"
// @Success 200 {array} dto.FundDTO
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /funds [get]
func (h *FundHandler) GetFunds(c *fiber.Ctx) error {
	var contributorID *uuid.UUID
	contributorIDStr := c.Query("contributorId")
	if contributorIDStr != "" {
		uuidVal, err := uuid.Parse(contributorIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contributor ID"})
		}
		contributorID = &uuidVal
	}

	if contributorID != nil {
		// Fetch funds contributed to (excluding managed funds)
		funds, err := h.service.GetContributedFunds(*contributorID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "Failed to fetch contributed funds"})
		}
		return c.Status(fiber.StatusOK).JSON(dto.MapFundsToDTOs(funds))
	}

	managerID := utils.GetCurrentUserID(c)
	if managerID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Fetch funds managed by the user
	funds, err := h.service.GetFundsManagedByUser(managerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch funds"})
	}
	return c.Status(fiber.StatusOK).JSON(dto.MapFundsToDTOs(funds))
}

// GetFunds retrieves all funds managed by the authenticated user
// @Summary Get all funds managed by the authenticated user
// @Description Returns a list of all funds where the authenticated user is the manager
// @Tags funds
// @Accept json
// @Produce json
// @Param fundId path string true "Fund ID"
// @Success 200 {object} dto.FundDTO
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /funds [get]
func (h *FundHandler) GetFund(c *fiber.Ctx) error {
	fundID, err := uuid.Parse(c.Params("fundId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Fund not found"})
	}

	fund, err := h.service.GetFund(fundID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch fund"})
	}

	if fund == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Fund not found"})
	}
	return c.Status(fiber.StatusOK).JSON(dto.MapFundToDTO(*fund))
}
