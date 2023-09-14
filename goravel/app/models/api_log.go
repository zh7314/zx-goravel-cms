package models

import (
	"goravel/app/utils/local"
)

const TableNameAPILog = "api_log"

// APILog mapped from table <api_log>
type APILog struct {
	ID           int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateAt     local.LocalTime `gorm:"column:create_at;not null;comment:创建时间" json:"create_at"`
	UpdateAt     local.LocalTime `gorm:"column:update_at;not null;comment:更新时间" json:"update_at"`
	Method       string          `gorm:"column:method;comment:请求方式" json:"method"`
	RequestIP    string          `gorm:"column:request_ip;comment:请求ip" json:"request_ip"`
	RequestURL   string          `gorm:"column:request_url;comment:请求url" json:"request_url"`
	QueryParams  string          `gorm:"column:query_params;comment:请求参数" json:"query_params"`
	ResponseData string          `gorm:"column:response_data;comment:返回的数据 不包含data" json:"response_data"`
}

// TableName APILog's table name
func (*APILog) TableName() string {
	return TableNameAPILog
}
