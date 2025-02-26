package repositories

import (
	"community-funds/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FundRepository struct {
	db *gorm.DB
}

func NewFundRepository(db *gorm.DB) *FundRepository {
	return &FundRepository{db}
}

// CreateFund adds a new fund to the database
func (r *FundRepository) CreateFund(fund *models.Fund) error {
	return r.db.Create(fund).Error
}

// GetFundByID retrieves a fund by ID
func (r *FundRepository) GetFundByID(id uuid.UUID) (*models.Fund, error) {
	var fund models.Fund
	err := r.db.Where("id = ?", id).First(&fund).Error
	return &fund, err
}

// GetFundsByManager retrieves all funds managed by a specific user
func (r *FundRepository) GetFundsByManager(managerID *uuid.UUID) ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Where("manager_id = ?", managerID).Find(&funds).Error
	return funds, err
}

// GetContributedFunds retrieves all funds a user has contributed to (excluding funds they manage)
func (r *FundRepository) GetContributedFunds(userID uuid.UUID) ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.
		Joins("Contributions", r.db.Where(&models.User{ID: userID})).
		Find(&funds).Error
	return funds, err
}
