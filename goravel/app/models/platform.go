package models

import "github.com/goravel/framework/support/carbon"

const TableNamePlatform = "platform"

type Platform struct {
	ID       int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	CreateAt *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Name     string           `gorm:"column:name" json:"name"`                           // comment 名称
	Sort     int              `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	UpdateAt *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Value    string           `gorm:"column:value" json:"value"`                         // comment 值

}

func (*Platform) TableName() string {
	return TableNamePlatform
}
