package services

import (
	"community-funds/pkg/models"
	"community-funds/pkg/repositories"

	"github.com/google/uuid"
)

type FundService struct {
	fundRepo *repositories.FundRepository
}

func NewFundService(fundRepo *repositories.FundRepository) *FundService {
	return &FundService{fundRepo}
}

func (s *FundService) CreateFund(name string, managerID uuid.UUID, targetAmount float64) (*models.Fund, error) {
	fund := &models.Fund{
		Name:         name,
		ManagerID:    managerID,
		TargetAmount: targetAmount,
	}
	err := s.fundRepo.CreateFund(fund)
	return fund, err
}

func (s *FundService) GetFund(fundID uuid.UUID) (*models.Fund, error) {
	return s.fundRepo.GetFundByID(fundID)
}

func (s *FundService) GetFundsByManagerID(managerID *uuid.UUID) ([]models.Fund, error) {
	return s.fundRepo.GetFundsByManagerID(managerID)
}

func (s *FundService) GetFundsByContributorID(contributorID uuid.UUID) ([]models.Fund, error) {
	return s.fundRepo.GetFundsByContributorID(contributorID)
}
