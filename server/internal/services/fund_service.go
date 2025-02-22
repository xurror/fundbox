package services

import (
	"community-funds/internal/models"
	"community-funds/internal/repositories"
	"errors"

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
func (s *FundService) GetFund(fundID uuid.UUID) (*models.Fund, error) {
	return s.fundRepo.GetFundByID(fundID)
}

// GetFundsManagedByUser retrieves all funds managed by a user
func (s *FundService) GetFundsManagedByUser(userID *uuid.UUID) ([]models.Fund, error) {
	return s.fundRepo.GetFundsByManager(userID)
}

// GetContributedFunds retrieves all funds a user has contributed to but does not manage
func (s *FundService) GetContributedFunds(userID uuid.UUID) ([]models.Fund, error) {
	funds, err := s.fundRepo.GetContributedFunds(userID)
	if err != nil {
		return nil, errors.New("failed to fetch contributed funds")
	}
	return funds, nil
}
