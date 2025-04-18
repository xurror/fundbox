package repositories

import (
	"community-funds/pkg/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// GetUserByAuth0ID finds a user by their Auth0 ID
func (r *UserRepository) GetUserByAuth0ID(auth0ID string) (*models.User, error) {
	var user models.User
	err := r.db.Where(&models.User{Auth0ID: auth0ID}).First(&user).Error
	return &user, err
}

// CreateUser inserts a new user into the database
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// GetUserByID retrieves a user by their internal ID
func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("FundsManaged").Preload("Contributions").First(&user, "id = ?", id).Error
	return &user, err
}
