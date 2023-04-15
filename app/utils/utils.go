package utils

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

func CheckPasswordHash(password string, userPassword string) bool {
	return true
}

func GenerateJWT(userId uuid.UUID) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString("privateKey")
}

func GetPageLimiAndOffset(ctx *gin.Context) (int, int) {
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
