package handlers

import (
	"net/http"

	"community-funds/internal/dto"
	"community-funds/internal/services"
	"community-funds/pkg/utils"

	"github.com/gin-gonic/gin"
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
func (h *ContributionHandler) CreateContribution(c *gin.Context) {
	var req struct {
		FundID uuid.UUID `json:"fundId" binding:"required"`
		Amount float64   `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contributorID := utils.GetCurrentUserID(c)
	anonymous := contributorID == nil

	contribution, err := h.service.MakeContribution(req.FundID, contributorID, req.Amount, anonymous)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contribute"})
		return
	}

	c.JSON(http.StatusCreated, dto.MapContributionToDTO(*contribution))
}

// GetContributions retrieves all contributions for a given fund or contributor or both
// @Summary Get contributions for a fund
// @Description Returns all contributions made to a specific fund, including contributor details
// @Tags contributions
// @Accept json
// @Produce json
// @Query fundId path string true "Fund ID"
// @Query contributorId path string true "Fund ID"
// @Success 200 {array} dto.ContributionDTO
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Server Error"
// @Security BearerAuth
// @Router /contributions [get]
func (h *ContributionHandler) GetContributions(c *gin.Context) {
	var fundID *uuid.UUID
	fundIDStr := c.Query("fundId")
	if fundIDStr != "" {
		uuidVal, err := uuid.Parse(fundIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fund ID"})
			return
		}
		fundID = &uuidVal
	}

	if fundID != nil {
		contributions, err := h.service.GetContributionsByFund(*fundID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contributions"})
			return
		}
		c.JSON(http.StatusOK, dto.MapContributionsToDTOs(contributions))
	} else {
		contributorID := utils.GetCurrentUserID(c)
		if contributorID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Contributor ID required"})
			return
		}
		contributions, err := h.service.GetContributionsByContributor(*contributorID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contributions"})
			return
		}
		c.JSON(http.StatusOK, dto.MapContributionsToDTOs(contributions))
	}
}
