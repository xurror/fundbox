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

func (r *FundRepository) GetFundsByManagerID(managerID *uuid.UUID) ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Where("manager_id = ?", managerID).Find(&funds).Error
	return funds, err
}

func (r *FundRepository) GetFundsByContributorID(contributorID uuid.UUID) ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Raw(`
	 SELECT * FROM funds
	 LEFT JOIN contributions ON funds.id = contributions.fund_id
	 WHERE contributions.contributor_id = ?
	`, contributorID).Scan(&funds).Error
	return funds, err
}
