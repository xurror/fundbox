package controllers

import (
	"getting-to-go/service"
	"getting-to-go/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FundController struct {
	fundService *services.FundService
}

func NewFundController(fundService *services.FundService) *FundController {
	return &FundController{
		fundService: fundService,
	}
}

func (c *FundController) Register(router *gin.RouterGroup) {
	router.GET("/funds", c.getFunds)
	router.GET("/funds/:id", c.getFund)
	router.POST("/funds", c.createFund)
	router.GET("/funds/:id/contributions", c.getFundContributions)
}

func (c *FundController) createFund(ctx *gin.Context) {
	var req struct {
		Reason      string `json:"reason" binding:"required"`
		Description string `json:"description"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}

	fund, err := c.fundService.CreateFund(req.Reason, req.Description)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, fund)
}

func (c *FundController) getFund(ctx *gin.Context) {
	id := ctx.Param("id")
	fund, err := c.fundService.GetFund(id)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, fund)
}

func (c *FundController) getFunds(ctx *gin.Context) {
	limit, offset := utils.GetPageLimitAndOffset(ctx)
	funds, err := c.fundService.GetFunds(limit, offset)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, funds)
}

func (c *FundController) getFundContributions(ctx *gin.Context) {
	id := ctx.Param("id")
	limit, offset := utils.GetPageLimitAndOffset(ctx)
	contributions, err := c.fundService.GetFundContributions(id, limit, offset)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, contributions)
}
