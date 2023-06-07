package services

import (
	"getting-to-go/models"
	"getting-to-go/utils"
	"net/http"
)

// UserService provides user-related services
type UserService struct{}

// NewUserService creates a new UserService instance
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(firstName, lastName, email, password string) (*models.User, error) {
	return models.CreateUser(&models.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Roles:     []string{"INITIATOR"},
	})
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id string) (*models.User, error) {
	return models.GetUser(id)
}

// GetUsers retrieves a list of users
func (s *UserService) GetUsers(limit, offset int) ([]*models.User, error) {
	return models.GetUsers(limit, offset)
}

// Authenticate authenticates a user by email and password
func (s *UserService) Authenticate(email, password string) (*models.User, error) {
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
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, utils.NewError(http.StatusUnauthorized, "Invalid email or password")
	}

	return user, nil
}
