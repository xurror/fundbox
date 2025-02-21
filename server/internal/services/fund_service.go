package services

import (
	"community-funds/internal/models"
	"community-funds/internal/repositories"

	"github.com/google/uuid"
)

type FundService struct {
	fundRepo *repositories.FundRepository
}

func NewFundService(fundRepo *repositories.FundRepository) *FundService {
	return &FundService{fundRepo}
}

// CreateFund associates a fund with a user (making them a fund manager)
func (s *FundService) CreateFund(name string, managerID uuid.UUID, targetAmount float64) (*models.Fund, error) {
	fund := &models.Fund{
		Name:         name,
		ManagerID:    managerID,
		TargetAmount: targetAmount,
	}
	err := s.fundRepo.CreateFund(fund)
	return fund, err
}

// GetFundsManagedByUser retrieves all funds managed by a user
func (s *FundService) GetFundsManagedByUser(userID *uuid.UUID) ([]models.Fund, error) {
	return s.fundRepo.GetFundsByManager(userID)
}
