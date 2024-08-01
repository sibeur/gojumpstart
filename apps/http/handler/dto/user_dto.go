package dto

type UserDTO struct {
	Username string `json:"username" validate:"required" v_msg:"required=Username harus diisi"`
	Email    string `json:"email" validate:"required,email" v_msg:"required=Email harus diisi;email=Email tidak valid"`
	Password string `json:"password" validate:"required" v_msg:"required=Password harus diisi"`
}
