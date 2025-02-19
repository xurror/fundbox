package repositories

import (
	"community-funds/internal/models"

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
func (r *FundRepository) GetFundByID(id string) (*models.Fund, error) {
	var fund models.Fund
	err := r.db.Where("id = ?", id).First(&fund).Error
	return &fund, err
}

// GetFundsByManager retrieves all funds managed by a specific user
func (r *FundRepository) GetFundsByManager(managerID string) ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Where("manager_id = ?", managerID).Find(&funds).Error
	return funds, err
}
