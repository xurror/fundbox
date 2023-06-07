package services

import (
	"getting-to-go/models"
	"github.com/google/uuid"
)

type ContributionService struct{}

func NewContributionService() *ContributionService {
	return &ContributionService{}
}

func (s *ContributionService) CreateContribution(fundID, contributorID uuid.UUID, amount float64, currencyID uuid.UUID) (*models.Contribution, error) {
	contribution := &models.Contribution{
		FundID:        fundID,
		ContributorID: contributorID,
	}

	contribution.Amount = models.Amount{Value: amount, CurrencyID: currencyID}
	return models.CreateContribution(contribution)
}

func (s *ContributionService) GetContribution(id string) (*models.Contribution, error) {
	return models.GetContribution(id)
}

func (s *ContributionService) GetContributions(limit, offset int) ([]*models.Contribution, error) {
	return models.GetContributions(limit, offset)
}
