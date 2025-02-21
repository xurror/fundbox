package repositories

import (
	"community-funds/internal/models"

	"github.com/google/uuid"
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
func (r *ContributionRepository) GetContributionsByFundOrContributor(fundID, contributorID *uuid.UUID) ([]models.Contribution, error) {
	var contributions []models.Contribution
	err := r.db.
		Preload("Contributor").
		Preload("Fund").
		Where("fund_id = ? OR contributor_id = ?", fundID, contributorID).
		Find(&contributions).Error
	return contributions, err
}
