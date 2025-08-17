package dto

type UserCreateDTO struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
	RoleId   int64  `json:"role_id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type UserEditDTO struct {
	Id       int64  `json:"id" binding:"required"`
	RoleId   int64  `json:"role_id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type UserListDTO struct {
	Pagination
	Username string `form:"username"`
	Phone    string `form:"keyword"`
	Email    string `form:"keyword"`
	Nickname string `form:"keyword"`
}
