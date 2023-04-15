package controllers

import (
	"getting-to-go/services"
	"getting-to-go/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContributorController struct {
	controllerService *services.ContributorService
}

func NewContributorController(controllerService *services.ContributorService) *ContributorController {
	return &ContributorController{controllerService: controllerService}
}

func (c *ContributorController) Register(router *gin.RouterGroup) {
	router.GET("/contributors", c.getContributors)
	router.GET("/contributors/:id", c.getContributor)
	router.POST("/contributors", c.createContributor)
	router.GET("/contributors/:id/contributions", c.getContributorsContributions)
}

func (c *ContributorController) createContributor(ctx *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.HandleBadRequest(ctx, err)
		return
	}

	contributor, err := c.controllerService.CreateContributor(req.Name)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, contributor)
}

func (c *ContributorController) getContributor(ctx *gin.Context) {
	id := ctx.Param("id")
	contributor, err := c.controllerService.GetContributor(id)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, contributor)
}

func (c *ContributorController) getContributors(ctx *gin.Context) {
	limit, offset := utils.GetPageLimiAndOffset(ctx)
	contributors, err := c.controllerService.GetContributors(limit, offset)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, contributors)
}

func (c *ContributorController) getContributorsContributions(ctx *gin.Context) {
	id := ctx.Param("id")
	limit, offset := utils.GetPageLimiAndOffset(ctx)
	contributions, err := c.controllerService.GetContributorsContributions(id, limit, offset)
	if err != nil {
		utils.HandleAppError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, contributions)
}
