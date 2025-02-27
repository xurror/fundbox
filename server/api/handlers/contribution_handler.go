package handlers

import (
	"community-funds/api/dto"
	"community-funds/pkg/services"
	"community-funds/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ContributionHandler handles contribution-related routes
type ContributionHandler struct {
	service *services.ContributionService
}

func NewContributionHandler(s *services.ContributionService) *ContributionHandler {
	return &ContributionHandler{service: s}
}

// CreateContribution handles contributions (Allows anonymous)
// @Summary Make a contribution
// @Description Anyone can contribute to a fund, authenticated users are tracked, anonymous users are allowed
// @Tags contributions
// @Accept json
// @Produce json
// @Success 201 {object} models.Contribution
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /contributions [post]
func (h *ContributionHandler) CreateContribution(c *fiber.Ctx) error {
	var req struct {
		FundID uuid.UUID `json:"fundId" binding:"required"`
		Amount float64   `json:"amount" binding:"required"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	contributorID := utils.GetCurrentUserID(c)
	anonymous := contributorID == nil

	contribution, err := h.service.MakeContribution(req.FundID, contributorID, req.Amount, anonymous)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to contribute"})
	}
	return c.Status(fiber.StatusCreated).JSON(dto.MapContributionToDTO(*contribution))
}

// GetContributions retrieves all contributions for a given fund or contributor or both
// @Summary Get contributions for a fund
// @Description Returns all contributions made to a specific fund, including contributor details
// @Tags contributions
// @Accept json
// @Produce json
// @Query fundId path string false "Fund ID"
// @Query contributorId path string false "Contributor ID"
// @Success 200 {array} dto.ContributionDTO
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Server Error"
// @Security BearerAuth
// @Router /contributions [get]
func (h *ContributionHandler) GetContributions(c *fiber.Ctx) error {
	var query struct {
		FundID        *uuid.UUID `query:"fundId"`
		ContributorID *uuid.UUID `query:"contributorId"`
	}

	_ = c.QueryParser(&query)

	if query.FundID != nil {
		contributions, err := h.service.GetContributionsByFund(*query.FundID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch contributions"})
		}

		return c.Status(fiber.StatusOK).JSON(dto.MapContributionsToDTOs(contributions))
	}

	if query.ContributorID != nil {
		contributions, err := h.service.GetContributionsByContributor(*query.ContributorID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch contributions"})
		}

		return c.Status(fiber.StatusOK).JSON(dto.MapContributionsToDTOs(contributions))
	}

	userID := utils.GetCurrentUserID(c)
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	contributions, err := h.service.GetManagedAndContributedContributions(*userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch contributions"})
	}

	return c.Status(fiber.StatusOK).JSON(dto.MapContributionsToDTOs(contributions))

}
