package service

import (
	models "getting-to-go/model"
	utils "getting-to-go/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthService struct {
	userService *UserService
}

// NewAuthService creates a new AuthService instance
func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{userService}
}

// Authenticate authenticates a user by email and password
func (s *AuthService) Authenticate(email, password string) (*models.User, error) {
	// Get the user by email
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// If the user does not exist, return an error
	if user == nil {
		return nil, utils.NewError(http.StatusUnauthorized, "Invalid email or password")
	}

	// If the password does not match, return an error
	if !utils.CheckPasswordHash(user.Password, password) {
		return nil, utils.NewError(http.StatusUnauthorized, "Invalid email or password")
	}

	return user, nil
}

func (s *AuthService) SignUp(c *gin.Context, firstName, lastName, email, password string, roles []string) (*models.User, error) {
	_, err := models.GetUserByEmail(email)
	if err == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email Already Exists"})
		log.Panic(err)
	}

	return s.userService.CreateUser(firstName, lastName, email, password, roles)
}
