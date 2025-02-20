package handlers

import (
	"net/http"

	"community-funds/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// FundHandler handles fund-related routes
type FundHandler struct {
	Logger  *logrus.Logger
	Service *services.FundService
}

func NewFundHandler(log *logrus.Logger, s *services.FundService) *FundHandler {
	return &FundHandler{
		Service: s,
		Logger:  log,
	}
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
	var req struct {
		Name         string  `json:"name" binding:"required"`
		TargetAmount float64 `json:"target_amount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var managerID uuid.UUID
	if id, exists := c.Get("user_id"); exists {
		managerID = id.(uuid.UUID)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	fund, err := h.Service.CreateFund(req.Name, managerID, req.TargetAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create fund"})
		return
	}

	c.JSON(http.StatusCreated, fund)
}
