package handlers

import (
	"community-funds/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FundHandler struct {
	Service *services.FundService
}

func NewFundHandler(s *services.FundService) *FundHandler {
	return &FundHandler{Service: s}
}

func (h *FundHandler) GetFunds(c *gin.Context) {
	funds := h.Service.GetAllFunds()
	c.JSON(http.StatusOK, gin.H{"funds": funds})
}
