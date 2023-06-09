package service

import (
	"getting-to-go/graph/generated"
	"getting-to-go/model"
	"getting-to-go/util"
	"github.com/google/uuid"
)

// UserService provides user-related services
type UserService struct{}

// NewUserService creates a new UserService instance
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(firstName, lastName, email, password string, roles []string) (*model.User, error) {
	r := model.FromStringArray(roles)
	return model.CreateUser(&model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  util.HashPassword(password),
		Roles:     model.ConvertToPQStringArray(r),
	})
}

func (s *UserService) CreateUserFromInput(input generated.NewUser) (*model.User, error) {
	return model.CreateUser(&model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  uuid.NewString(),
		Roles:     model.ConvertToPQStringArray(input.Roles),
	})
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(id uuid.UUID) (*model.User, error) {
	return model.GetUser(id)
}

// GetUsers retrieves a list of users
func (s *UserService) GetUsers(limit, offset int) ([]*model.User, error) {
	return model.GetUsers(limit, offset)
}
