package services

import (
	"community-funds/internal/models"
	"community-funds/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) RegisterUser(name, email, auth0ID, role string) (*models.User, error) {
	user := &models.User{Name: name, Email: email, Auth0ID: auth0ID}
	err := s.repo.CreateUser(user)
	return user, err
}
