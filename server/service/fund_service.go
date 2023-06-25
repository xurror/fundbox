package service

import (
	"getting-to-go/model"
	"gorm.io/gorm"
)

type FundService struct {
	db *gorm.DB
}

func NewFundService(db *gorm.DB) *FundService {
	return &FundService{
		db: db,
	}
}

func (s *FundService) CreateFund(reason, description string) (*model.Fund, error) {
	fund := &model.Fund{
		Reason:      reason,
		Description: description,
	}
	result := s.db.Create(&fund)
	return fund, model.HandleError(result.Error)
}

func (s *FundService) GetFund(id string) (*model.Fund, error) {
	fund := &model.Fund{}
	result := s.db.First(&fund, "id = ?", id)
	return fund, model.HandleError(result.Error)
}

func (s *FundService) GetFunds(limit, offset int) ([]*model.Fund, error) {
	var funds []*model.Fund
	result := s.db.Limit(limit).Offset(offset).Find(&funds)
	return funds, model.HandleError(result.Error)
}

func (s *FundService) GetFundContributions(fundID string, limit, offset int) ([]*model.Contribution, error) {
	var contributions []*model.Contribution
	result := s.db.Limit(limit).
		Offset(offset).
		Find(&contributions, "fund_id = ?", fundID)
	return contributions, model.HandleError(result.Error)
}
