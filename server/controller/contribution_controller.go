package controller

import (
	"getting-to-go/service"
	"getting-to-go/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContributionController struct {
	contributionService *service.ContributionService
}

func NewContributionController(contributionService *service.ContributionService) *ContributionController {
	return &ContributionController{
		contributionService: contributionService,
	}
}

func (c *ContributionController) Register(router *gin.RouterGroup) {
	router.GET("/contributions", c.getContributions)
	router.GET("/contributions/:id", c.getContribution)
	router.POST("/contributions", c.createContribution)
}

func (c *ContributionController) createContribution(ctx *gin.Context) {
	var req struct {
		FundID        uuid.UUID `json:"fund_id" binding:"required"`
		ContributorID uuid.UUID `json:"contributor_id" binding:"required"`
		CurrencyID    uuid.UUID `json:"currency_id" binding:"required"`
		Amount        float64   `json:"amount" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleBadRequest(ctx, err)
		return
	}

	contribution, err := c.contributionService.CreateContribution(
		req.FundID,
		req.ContributorID,
		req.Amount,
		req.CurrencyID,
	)
	if err != nil {
		util.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, contribution)
}

func (c *ContributionController) getContribution(ctx *gin.Context) {
	id := ctx.Param("id")
	contribution, err := c.contributionService.GetContribution(id)
	if err != nil {
		util.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, contribution)
}

func (c *ContributionController) getContributions(ctx *gin.Context) {
	limit, offset := util.GetPageLimitAndOffset(ctx)
	contributions, err := c.contributionService.GetContributions(limit, offset)
	if err != nil {
		util.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, contributions)
}