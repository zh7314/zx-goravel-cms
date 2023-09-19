package models

import (
	"goravel/app/utils/local"
)

const TableNameFeedback = "feedback"

type Feedback struct {
	ID       int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	Contact  string          `gorm:"column:contact" json:"contact"`                     // comment 联系方式
	Content  string          `gorm:"column:content" json:"content"`                     // comment 反馈内容
	CreateAt local.LocalTime `gorm:"-" json:"create_at"`                                // comment 创建时间
	Lang     string          `gorm:"column:lang" json:"lang"`                           // comment 语言类型
	NickName string          `gorm:"column:nick_name" json:"nick_name"`                 // comment 反馈昵称
	Platform string          `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	UpdateAt local.LocalTime `gorm:"-" json:"update_at"`                                // comment 更新时间

}

func (*Feedback) TableName() string {
	return TableNameFeedback
}
