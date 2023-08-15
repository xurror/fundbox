package service

import (
	"getting-to-go/graph/generated"
	"getting-to-go/model"
	"getting-to-go/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// UserService provides user-related services
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new UserService instance
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(firstName, lastName, email, password string, roles []model.Role) (*model.User, error) {
	user := &model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  util.HashPassword(password),
		Roles:     model.ConvertToPQStringArray(roles),
	}
	result := s.db.Create(&user)
	return user, model.HandleError(result.Error)
}

func (s *UserService) CreateUserFromInput(input generated.NewUser) (*model.User, error) {
	return s.CreateUser(
		input.FirstName,
		input.LastName,
		input.Email,
		uuid.NewString(),
		input.Roles,
	)
}

func (s *UserService) GetUser(id uuid.UUID) (*model.User, error) {
	user := &model.User{}
	result := s.db.Preload(clause.Associations).First(&user, id)
	return user, model.HandleError(result.Error)
}

func (s *UserService) GetUsers(limit, offset int) ([]*model.User, error) {
	var users []*model.User
	result := s.db.Preload(clause.Associations).Limit(limit).Offset(offset).Find(&users)
	return users, model.HandleError(result.Error)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	result := s.db.Preload(clause.Associations).First(&user, "email = ?", email)
	return user, model.HandleError(result.Error)
}

func (s *UserService) GetUserContributions(userID uuid.UUID, limit, offset int) ([]*model.Contribution, error) {
	var contributions []*model.Contribution
	result := s.db.
		Preload(clause.Associations).
		Limit(limit).
		Offset(offset).
		Find(&contributions, "contributor_id = ?", userID)
	return contributions, model.HandleError(result.Error)
}
