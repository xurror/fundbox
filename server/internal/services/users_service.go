package services

import "community-funds/internal/repositories"

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(r *repositories.UserRepository) *UserService {
	return &UserService{Repo: r}
}

func (s *UserService) GetAllUsers() []string {
	return s.Repo.GetAll()
}
