package repositories

import (
	"community-funds/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContributionRepository struct {
	db *gorm.DB
}

func NewContributionRepository(db *gorm.DB) *ContributionRepository {
	return &ContributionRepository{db}
}

func (r *ContributionRepository) CreateContribution(contribution *models.Contribution) error {
	return r.db.Create(contribution).Error
}

func (r *ContributionRepository) GetContributionsByFund(fundID uuid.UUID) ([]models.Contribution, error) {
	var contributions []models.Contribution
	err := r.db.
		Preload("Contributor").
		Preload("Fund").
		Where("fund_id = ? ", fundID).
		Find(&contributions).Error
	return contributions, err
}

func (r *ContributionRepository) GetContributionsByContributor(contributorID uuid.UUID) ([]models.Contribution, error) {
	var contributions []models.Contribution
	err := r.db.
		Preload("Contributor").
		Preload("Fund").
		Where("contributor_id = ?", contributorID).
		Find(&contributions).Error
	return contributions, err
}

func (r *ContributionRepository) GetFundsByContributorIDOrManageID(userID uuid.UUID) ([]models.Contribution, error) {
	var contributions []models.Contribution
	err := r.db.
		Preload("Contributor").
		Preload("Fund").
		Where(&models.Contribution{
			ContributorID: &userID,
		}).
		Or(&models.Contribution{
			Fund: models.Fund{
				ManagerID: userID,
			},
		}).
		Find(&contributions).Error
	return contributions, err
}
