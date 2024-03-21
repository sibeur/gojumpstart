package repository

import (
	"gojumpstart/core/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FindAll() ([]*entity.User, error) {
	var users []*entity.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) Create(user *entity.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
