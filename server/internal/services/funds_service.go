package services

import "community-funds/internal/repositories"

type FundService struct {
	Repo *repositories.FundRepository
}

func NewFundService(r *repositories.FundRepository) *FundService {
	return &FundService{Repo: r}
}

func (s *FundService) GetAllFunds() []string {
	return s.Repo.GetAll()
}
