package service

import (
	"getting-to-go/model"
	"github.com/google/uuid"
)

type ContributionService struct{}

func NewContributionService() *ContributionService {
	return &ContributionService{}
}

func (s *ContributionService) CreateContribution(fundID, contributorID uuid.UUID, amount float64, currencyID uuid.UUID) (*model.Contribution, error) {
	contribution := &model.Contribution{
		FundID:        fundID,
		ContributorID: contributorID,
	}

	contribution.Amount = model.Amount{Value: amount, CurrencyID: currencyID}
	return model.CreateContribution(contribution)
}

func (s *ContributionService) GetContribution(id string) (*model.Contribution, error) {
	return model.GetContribution(id)
}

func (s *ContributionService) GetContributions(limit, offset int) ([]*model.Contribution, error) {
	return model.GetContributions(limit, offset)
}

func (s *ContributionService) GetContributionsByUserID(userId uuid.UUID, limit, offset int) ([]*model.Contribution, error) {
	return model.GetUserContributions(userId, limit, offset)
}
