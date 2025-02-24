package services

import (
	"community-funds/pkg/models"
	"community-funds/pkg/repositories"

	"github.com/google/uuid"
)

type ContributionService struct {
	repo *repositories.ContributionRepository
}

func NewContributionService(repo *repositories.ContributionRepository) *ContributionService {
	return &ContributionService{repo}
}

func (s *ContributionService) MakeContribution(fundID uuid.UUID, contributorID *uuid.UUID, amount float64, anonymous bool) (*models.Contribution, error) {

	contribution := &models.Contribution{
		FundID:        fundID,
		ContributorID: contributorID,
		Amount:        amount,
		Anonymous:     anonymous,
	}
	err := s.repo.CreateContribution(contribution)
	return contribution, err
}

func (s *ContributionService) GetContributionsByFund(fundID uuid.UUID) ([]models.Contribution, error) {
	contributions, err := s.repo.GetContributionsByFund(fundID)
	if err != nil {
		return nil, err
	}
	return contributions, nil
}

func (s *ContributionService) GetContributionsByContributor(contributorID uuid.UUID) ([]models.Contribution, error) {
	contributions, err := s.repo.GetContributionsByContributor(contributorID)
	if err != nil {
		return nil, err
	}
	return contributions, nil
}
