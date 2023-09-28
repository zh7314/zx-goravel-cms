package models

import "github.com/goravel/framework/support/carbon"

const TableNameLang = "lang"

type Lang struct {
	ID       int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	CreateAt *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Name     string           `gorm:"column:name" json:"name"`                           // comment 语言名称
	Sort     int              `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	UpdateAt *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Value    string           `gorm:"column:value" json:"value"`                         // comment 语言值

}

func (*Lang) TableName() string {
	return TableNameLang
}
