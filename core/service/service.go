package service

import "gojumpstart/core/repository"

type Service struct {
	User *UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
	}
}
