package handlers

import (
	"net/http"

	"community-funds/internal/services"

	"github.com/gin-gonic/gin"
)

// FundHandler handles fund-related routes
type FundHandler struct {
	Service *services.FundService
}

func NewFundHandler(s *services.FundService) *FundHandler {
	return &FundHandler{Service: s}
}

// CreateFund handles fund creation (Only Fund Managers)
// @Summary Create a fund
// @Description Fund managers can create a fund with a target amount
// @Tags funds
// @Accept json
// @Produce json
// @Success 201 {object} models.Fund
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /funds [post]
func (h *FundHandler) CreateFund(c *gin.Context) {
	userRole := c.GetString("user_role")
	if userRole != "fund_manager" && userRole != "both" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only fund managers can create funds"})
		return
	}

	var req struct {
		Name         string  `json:"name" binding:"required"`
		TargetAmount float64 `json:"target_amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	managerID := c.GetString("user_id")
	fund, err := h.Service.CreateFund(req.Name, managerID, req.TargetAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create fund"})
		return
	}

	c.JSON(http.StatusCreated, fund)
}
