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
	return u.repo.User.FindAll()
}

func (u *UserService) Create(user *entity.User) error {
	return u.repo.User.Create(user)
}
