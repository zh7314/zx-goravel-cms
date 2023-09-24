package models

import "github.com/goravel/framework/support/carbon"

const TableNameConfig = "config"

type Config struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    Type            string           `gorm:"column:type" json:"type"`           // comment 标签/类型
    Name            string           `gorm:"column:name" json:"name"`           // comment 名称
    Value            string           `gorm:"column:value" json:"value"`           // comment 值
    CreateAt            carbon.DateTime           `gorm:"column:create_at;->" json:"create_at"`           // comment 创建时间
    UpdateAt            carbon.DateTime           `gorm:"column:update_at;->" json:"update_at"`           // comment 更新时间

}

func (*Config) TableName() string {
	return TableNameConfig
}
