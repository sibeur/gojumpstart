package entity

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) ToJSON() map[string]any {
	return map[string]any{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
	}
}
