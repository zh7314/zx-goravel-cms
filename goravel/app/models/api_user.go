package models

import (
	"goravel/app/utils/local"
	"gorm.io/gorm"
)

const TableNameAPIUser = "api_user"

// APIUser mapped from table <api_user>
type APIUser struct {
	ID         int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	APIKey     string          `gorm:"column:api_key;not null;comment:接口用户名" json:"api_key"`
	APISecret  string          `gorm:"column:api_secret;not null;comment:接口秘钥" json:"api_secret"`
	CreateTime local.LocalTime `gorm:"column:create_time;not null;comment:创建时间" json:"create_time"`
	UpdateTime local.LocalTime `gorm:"column:update_time;not null;comment:更新时间" json:"update_time"`
	IPList     string          `gorm:"column:ip_list;comment:api接口访问ip列表" json:"ip_list"`
	Status     int             `gorm:"column:status;not null;default:10;comment:10默认未启用20启用30禁止" json:"status"`
	EndDate    local.LocalTime `gorm:"column:end_date;comment:api终止使用日期" json:"end_date"`
	DeletedAt  gorm.DeletedAt  `gorm:"column:deleted_at;comment:是否删除" json:"deleted_at"`
}

// TableName APIUser's table name
func (*APIUser) TableName() string {
	return TableNameAPIUser
}
