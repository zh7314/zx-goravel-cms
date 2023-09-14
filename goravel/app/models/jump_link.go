package models

import (
	"goravel/app/utils/local"
	"gorm.io/gorm"
)

const TableNameJumpLink = "jump_link"

// JumpLink mapped from table <jump_link>
//type JumpLink struct {
//	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
//	JumpURL   string         `gorm:"column:jump_url;comment:需要跳转到url" json:"jump_url"`
//	HashKey   string         `gorm:"column:hash_key;comment:id的哈希值" json:"hash_key"`
//	CreatedAt time.Time      `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
//	UpdatedAt time.Time      `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`
//	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:是否删除" json:"deleted_at"`
//	EndTime   time.Time      `gorm:"column:end_time;comment:最后存活时间" json:"end_time"`
//}

// JumpLink mapped from table <jump_link>
type JumpLink struct {
	ID        int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	JumpURL   string          `gorm:"column:jump_url;comment:需要跳转到url" json:"jump_url"`
	HashKey   string          `gorm:"column:hash_key;comment:id的哈希值" json:"hash_key"`
	CreatedAt local.LocalTime `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`
	UpdatedAt local.LocalTime `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt  `gorm:"column:deleted_at;comment:是否删除" json:"deleted_at"`
	EndTime   local.LocalTime `gorm:"column:end_time;comment:最后存活时间" json:"end_time"`
}

// TableName JumpLink's table name
func (*JumpLink) TableName() string {
	return TableNameJumpLink
}
