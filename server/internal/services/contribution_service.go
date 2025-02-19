package services

import (
	"community-funds/internal/models"
	"community-funds/internal/repositories"

	"github.com/google/uuid"
)

type ContributionService struct {
	repo *repositories.ContributionRepository
}

func NewContributionService(repo *repositories.ContributionRepository) *ContributionService {
	return &ContributionService{repo}
}

func (s *ContributionService) MakeContribution(fundID, contributorID string, amount float64, anonymous bool) (*models.Contribution, error) {
	fundUUID, err := uuid.Parse(fundID)
	if err != nil {
		return nil, err
	}
	contributorUUID, err := uuid.Parse(contributorID)
	if err != nil {
		return nil, err
	}
	contribution := &models.Contribution{FundID: fundUUID, ContributorID: &contributorUUID, Amount: amount, Anonymous: anonymous}
	err = s.repo.CreateContribution(contribution)
	return contribution, err
}
