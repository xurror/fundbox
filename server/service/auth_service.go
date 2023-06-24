package service

import (
	models "getting-to-go/model"
	_type "getting-to-go/type"
	utils "getting-to-go/util"
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

	if user == nil {
		return nil, _type.NewErrorResponse(http.StatusUnauthorized, "invalid email or password")
	}

	if !utils.CheckPasswordHash(user.Password, password) {
		return nil, _type.NewErrorResponse(http.StatusUnauthorized, "invalid email or password")
	}

	return user, nil
}

func (s *AuthService) SignUp(firstName, lastName, email, password string, roles []string) (*models.User, error) {
	_, err := models.GetUserByEmail(email)
	if err == nil {
		log.Print(err)
		return nil, _type.NewErrorResponse(http.StatusConflict, "email already exists")
	}

	return s.userService.CreateUser(firstName, lastName, email, password, roles)
}
