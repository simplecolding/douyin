// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLove = "love"

// Love mapped from table <love>
type Love struct {
	Lid       int64     `gorm:"column:lid;primaryKey;autoIncrement:true" json:"lid"`           // 主键
	Vid       int64     `gorm:"column:vid;not null" json:"vid"`                                // 视频id
	UID       int64     `gorm:"column:uid;not null" json:"uid"`                                // 用户id
	Status    bool      `gorm:"column:status;default:1" json:"status"`                         // 是否删除 1:是  0:否
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"` // 更新时间
}

// TableName Love's table name
func (*Love) TableName() string {
	return TableNameLove
}
