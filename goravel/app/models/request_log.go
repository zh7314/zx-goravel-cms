package models

import (
	"github.com/goravel/framework/support/carbon"
	"time"
)

const TableNameRequestLog = "request_log"

type RequestLog struct {
	ID       int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	CreateAt *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Data     string           `gorm:"column:data" json:"data"`                           // comment 返回的数据 不包含data
	Header   string           `gorm:"column:header" json:"header"`                       // comment 请求header
	Ip       string           `gorm:"column:ip" json:"ip"`                               // comment 请求ip
	Method   string           `gorm:"column:method" json:"method"`                       // comment 请求方式
	Params   string           `gorm:"column:params" json:"params"`                       // comment 请求参数
	ReturnAt *time.Time       `gorm:"column:return_at" json:"return_at"`                 // comment 返回结果的时间
	UpdateAt *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Url      string           `gorm:"column:url" json:"url"`                             // comment 请求url

}

func (*RequestLog) TableName() string {
	return TableNameRequestLog
}
