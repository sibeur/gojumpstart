package repository

import (
	"gojumpstart/core/common/helper"
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

func (u *UserRepository) FindAllPaginate(currentPage, perPage int64, filters *entity.UserListFilter) (*helper.Pagination, error) {

	var users []*entity.User
	query := u.db
	query = filters.FilterQuery(query)
	offset := (currentPage - 1) * perPage
	query = query.Offset(int(offset))
	query = query.Limit(int(perPage))
	err := query.Model(&entity.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}

	queryCount := u.db
	queryCount = filters.FilterQuery(queryCount)
	var total int64
	err = queryCount.Model(&entity.User{}).Count(&total).Error
	if err != nil {
		return nil, err
	}
	meta := helper.NewPaginationMeta(currentPage, perPage, total)
	pagination := helper.NewPagination(users, meta)
	return pagination, nil
}

func (u *UserRepository) Create(user *entity.User) error {
	err := u.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
