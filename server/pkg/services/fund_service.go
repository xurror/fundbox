package services

import (
	"community-funds/pkg/models"
	"community-funds/pkg/repositories"

	"github.com/google/uuid"
)

type FundService struct {
	stripeService *StripeService
	fundRepo      *repositories.FundRepository
}

func NewFundService(fundRepo *repositories.FundRepository, stripeService *StripeService) *FundService {
	return &FundService{
		fundRepo:      fundRepo,
		stripeService: stripeService,
	}
}

func (s *FundService) CreateFund(name string, managerID uuid.UUID, targetAmount float64) (*models.Fund, error) {
	accountID, err := s.stripeService.CreateAccount()
	if err != nil {
		return nil, err
	}

	fund := &models.Fund{
		Name:                     name,
		ManagerID:                managerID,
		TargetAmount:             targetAmount,
		StripeConnectedAccountId: *accountID,
	}
	err = s.fundRepo.CreateFund(fund)
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
