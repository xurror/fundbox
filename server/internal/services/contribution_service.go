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
