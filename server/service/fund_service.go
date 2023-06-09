package service

import "getting-to-go/model"

type FundService struct{}

func NewFundService() *FundService {
	return &FundService{}
}

func (s *FundService) CreateFund(reason, description string) (*model.Fund, error) {
	return model.CreateFund(&model.Fund{
		Reason:      reason,
		Description: description,
	})
}

func (s *FundService) GetFund(id string) (*model.Fund, error) {
	return model.GetFund(id)
}

func (s *FundService) GetFunds(limit, offset int) ([]*model.Fund, error) {
	return model.GetFunds(limit, offset)
}

func (s *FundService) GetFundContributions(fundID string, limit, offset int) ([]*model.Contribution, error) {
	return model.GetFundContributions(fundID, limit, offset)
}
