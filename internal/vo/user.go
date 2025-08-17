package vo

type UserInfoVO struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	RoleId   int64  `json:"role_id"`
	Ctime    int64  `json:"create_time"`
}

type UserListVo struct {
	Data  []UserInfoVO `json:"data"`
	Total int64        `json:"total"`
}
