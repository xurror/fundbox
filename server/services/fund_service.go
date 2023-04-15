package services

import (
	"getting-to-go/models"
)

type FundService struct{}

func NewFundService() *FundService {
	return &FundService{}
}

func (s *FundService) CreateFund(reason, description string) (*models.Fund, error) {
	return models.CreateFund(&models.Fund{
		Reason:      reason,
		Description: description,
	})
}

func (s *FundService) GetFund(id string) (*models.Fund, error) {
	return models.GetFund(id)
}

func (s *FundService) GetFunds(limit, offset int) ([]*models.Fund, error) {
	return models.GetFunds(limit, offset)
}

func (s *FundService) GetFundContributions(fundID string, limit, offset int) ([]*models.Contribution, error) {
	return models.GetFundContributions(fundID, limit, offset)
}
