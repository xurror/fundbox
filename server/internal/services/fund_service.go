package services

import (
	"community-funds/internal/models"
	"community-funds/internal/repositories"

	"github.com/google/uuid"
)

type FundService struct {
	repo *repositories.FundRepository
}

func NewFundService(repo *repositories.FundRepository) *FundService {
	return &FundService{repo}
}

func (s *FundService) CreateFund(name string, managerID string, targetAmount float64) (*models.Fund, error) {
	managerUUID, err := uuid.Parse(managerID)
	if err != nil {
		return nil, err
	}

	fund := &models.Fund{Name: name, ManagerID: managerUUID, TargetAmount: targetAmount}
	err = s.repo.CreateFund(fund)
	return fund, err
}
