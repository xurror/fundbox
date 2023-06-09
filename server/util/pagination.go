package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPageLimitAndOffset(ctx *gin.Context) (int, int) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		HandleBadRequest(ctx, err)
		panic(err)
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "1"))
	if err != nil {
		HandleBadRequest(ctx, err)
		panic(err)
	}
	return limit, offset
}

func GetLimitAndOffset(limit *int, offset *int) (int, int) {
	var l = 10
	if limit != nil {
		l = *limit
	}

	var o = 10
	if offset != nil {
		o = *offset
	}
	return l, o
}
