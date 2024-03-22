package service

import "gojumpstart/core/repository"

type Service struct {
	User *UserService
	Todo *TodoService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
		Todo: NewTodoService(repo),
	}
}
