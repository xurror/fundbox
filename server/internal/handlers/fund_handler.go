package handlers

import (
	"net/http"

	"community-funds/internal/dto"
	"community-funds/internal/services"
	"community-funds/pkg/utils"

	"github.com/gin-gonic/gin"
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
func (h *FundHandler) CreateFund(c *gin.Context) {
	var req struct {
		Name         string  `json:"name" binding:"required"`
		TargetAmount float64 `json:"targetAmount" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	managerID := utils.GetCurrentUserID(c)
	fund, err := h.Service.CreateFund(req.Name, *managerID, req.TargetAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create fund"})
		return
	}

	c.JSON(http.StatusCreated, dto.MapFundToDTO(*fund))
}

// GetFunds retrieves all funds managed by the authenticated user
// @Summary Get all funds managed by the authenticated user
// @Description Returns a list of all funds where the authenticated user is the manager
// @Tags funds
// @Accept json
// @Produce json
// @Success 200 {array} dto.FundDTO
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /funds [get]
func (h *FundHandler) GetFunds(c *gin.Context) {
	managerID := utils.GetCurrentUserID(c)

	// Fetch funds managed by the user
	funds, err := h.Service.GetFundsManagedByUser(managerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch funds"})
		return
	}

	c.JSON(http.StatusOK, dto.MapFundsToDTOs(funds))
}

// GetContributedFunds retrieves all funds a user has contributed to (excluding those they manage)
// @Summary Get funds contributed to
// @Description Returns a list of all funds a user has contributed to but does not manage
// @Tags funds
// @Accept json
// @Produce json
// @Success 200 {array} dto.FundDTO
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Server error"
// @Security BearerAuth
// @Router /funds/contributed [get]
func (h *FundHandler) GetContributedFunds(c *gin.Context) {
	userID := utils.GetCurrentUserID(c)

	// Fetch funds contributed to (excluding managed funds)
	funds, err := h.Service.GetContributedFunds(*userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch contributed funds"})
		return
	}

	c.JSON(http.StatusOK, dto.MapFundsToDTOs(funds))
}
