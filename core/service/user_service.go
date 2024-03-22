package service

import (
	"gojumpstart/core/entity"
	"gojumpstart/core/repository"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) FindAll() ([]*entity.User, error) {
	result, err := u.repo.User.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserService) Create(user *entity.User) error {
	return u.repo.User.Create(user)
}
