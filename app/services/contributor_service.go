package services

import (
	"getting-to-go/app/models"
)

type ContributorService struct{}

func NewContributorService() *ContributorService {
	return &ContributorService{}
}

func (s *ContributorService) CreateContributor(name string) (*models.Contributor, error) {
	return models.CreateContributor(&models.Contributor{
		Name: name,
	})
}

func (s *ContributorService) GetContributor(id string) (*models.Contributor, error) {
	return models.GetContributor(id)
}

func (s *ContributorService) GetContributors(limit, offset int) ([]*models.Contributor, error) {
	return models.GetContributors(limit, offset)
}

func (s *ContributorService) GetContributorsContributions(contributorID string, limit, offset int) ([]*models.Contribution, error) {
	return models.GetContributorContributions(contributorID, limit, offset)
}
