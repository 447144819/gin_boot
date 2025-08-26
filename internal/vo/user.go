package vo

type UserInfoVO struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	RoleId   int64  `json:"role_id"`
	Ctime    int64  `json:"create_time"`
}
