package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func HandleBadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": err.Error(),
	})
}

func HandleError(ctx *gin.Context, code int, message string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": message,
	})
}

func HandleAppError(ctx *gin.Context, err error) {
	if appError, ok := err.(*AppError); ok {
		ctx.JSON(appError.Code, gin.H{
			"message": appError.Message,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
