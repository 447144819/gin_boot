package dto

type UserCreateDTO struct {
	Username string `json:"username" binding:"required,min=3,max=20" label:"用户名"`
	Password string `json:"password" binding:"required,min=6,max=20" validate:"required,min=6,max=20" label:"密码"`
	RoleId   int64  `json:"role_id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
