package handlers

import (
	"net/http"

	"community-funds/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ContributionHandler handles contribution-related routes
type ContributionHandler struct {
	Service *services.ContributionService
}

func NewContributionHandler(s *services.ContributionService) *ContributionHandler {
	return &ContributionHandler{Service: s}
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
		FundID uuid.UUID `json:"fund_id" binding:"required"`
		Amount float64   `json:"amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var anonymous bool
	var contributorID *uuid.UUID

	if id, exists := c.Get("user_id"); exists {
		contributorID = id.(*uuid.UUID)
		anonymous = false
	} else {
		anonymous = true
	}

	contribution, err := h.Service.MakeContribution(req.FundID, contributorID, req.Amount, anonymous)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to contribute"})
		return
	}

	c.JSON(http.StatusCreated, contribution)
}
