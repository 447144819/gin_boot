package model

import "gorm.io/plugin/soft_delete"

type CommonModel struct {
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag;default:0;comment:0未删除 1已删除"`
	CreateTime int64                 `gorm:"autoCreateTime"`
	UpdateTime int64                 `gorm:"autoUpdateTime"`
}
