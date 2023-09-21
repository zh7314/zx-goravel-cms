package models

import "goravel/app/utils/local"

const TableNameRequestLog = "request_log"

type RequestLog struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    CreateAt            local.LocalTime           `gorm:"-" json:"create_at"`           // comment 创建时间
    UpdateAt            local.LocalTime           `gorm:"-" json:"update_at"`           // comment 更新时间
    Method            string           `gorm:"column:method" json:"method"`           // comment 请求方式
    Ip            string           `gorm:"column:ip" json:"ip"`           // comment 请求ip
    Url            string           `gorm:"column:url" json:"url"`           // comment 请求url
    Params            string           `gorm:"column:params" json:"params"`           // comment 请求参数
    Header            string           `gorm:"column:header" json:"header"`           // comment 请求header
    Data            string           `gorm:"column:data" json:"data"`           // comment 返回的数据 不包含data
    ReturnAt            local.LocalTime           `gorm:"column:return_at" json:"return_at"`           // comment 返回结果的时间

}

func (*RequestLog) TableName() string {
	return TableNameRequestLog
}
