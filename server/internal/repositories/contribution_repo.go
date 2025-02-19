package repositories

import (
	"community-funds/internal/models"

	"gorm.io/gorm"
)

type ContributionRepository struct {
	db *gorm.DB
}

func NewContributionRepository(db *gorm.DB) *ContributionRepository {
	return &ContributionRepository{db}
}

// CreateContribution records a contribution to a fund
func (r *ContributionRepository) CreateContribution(contribution *models.Contribution) error {
	return r.db.Create(contribution).Error
}

// GetContributionsByFund retrieves all contributions for a given fund
func (r *ContributionRepository) GetContributionsByFund(fundID string) ([]models.Contribution, error) {
	var contributions []models.Contribution
	err := r.db.Where("fund_id = ?", fundID).Find(&contributions).Error
	return contributions, err
}

// GetContributionsByUser retrieves all contributions made by a specific user
func (r *ContributionRepository) GetContributionsByUser(userID string) ([]models.Contribution, error) {
	var contributions []models.Contribution
	err := r.db.Where("contributor_id = ?", userID).Find(&contributions).Error
	return contributions, err
}
