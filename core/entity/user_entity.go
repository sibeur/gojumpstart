package entity

import "gorm.io/gorm"

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserListFilter struct {
	Search string
}

func (f *UserListFilter) FilterQuery(query *gorm.DB) *gorm.DB {
	if f.Search != "" {
		query = query.Where("username LIKE ? OR email LIKE ?", "%"+f.Search+"%", "%"+f.Search+"%")
	}

	return query
}

func (u *User) ToJSON() map[string]any {
	return map[string]any{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
	}
}
