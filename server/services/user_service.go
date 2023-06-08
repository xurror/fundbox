package services

import (
	"getting-to-go/graph/generated"
	"getting-to-go/models"
	"getting-to-go/utils"
	"github.com/google/uuid"
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
		Roles:     models.ConvertToPQStringArray([]models.Role{"INITIATOR"}),
	})
}

func (s *UserService) CreateUserFromInput(input generated.NewUser) (*models.User, error) {
	return models.CreateUser(&models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  uuid.NewString(),
		Roles:     models.ConvertToPQStringArray(input.Roles),
	})
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id uuid.UUID) (*models.User, error) {
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
