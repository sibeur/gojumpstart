package service

import (
	"gojumpstart/core/entity"
	"gojumpstart/core/repository"
)

type TodoService struct {
	repo *repository.Repository
}

func NewTodoService(repo *repository.Repository) *TodoService {
	return &TodoService{repo: repo}
}

func (u *TodoService) FindAll() ([]*entity.Todo, error) {
	result, err := u.repo.Todo.FindAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *TodoService) Create(user *entity.Todo) error {
	return u.repo.Todo.Create(user)
}
