package model

import "gorm.io/plugin/soft_delete"

type User struct {
	Id       uint64                `gorm:"primary_key;auto_increment"`
	Username string                `gorm:"unique;type:varchar(30);not null;comment:用户名"`
	Password string                `gorm:"type:varchar(200);not null;comment:密码"`
	Nickname string                `gorm:"type:varchar(100);comment:用户昵称"`
	Phone    string                `gorm:"type:char(11);comment:手机号"`
	Email    string                `gorm:"type:varchar(200);comment:邮箱"`
	RoleId   int64                 `gorm:"type:int;comment:角色id"`
	IsDel    soft_delete.DeletedAt `gorm:"softDelete:flag;comment:0未删除 1已删除"`
	Ctime    int64                 `gorm:"autoCreateTime:milli"`
	Utime    int64                 `gorm:"autoUpdateTime:milli"`
}
