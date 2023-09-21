package models

import "goravel/app/utils/local"

const TableNameConfig = "config"

type Config struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    Type            string           `gorm:"column:type" json:"type"`           // comment 标签/类型
    Name            string           `gorm:"column:name" json:"name"`           // comment 名称
    Value            string           `gorm:"column:value" json:"value"`           // comment 值
    CreateAt            local.LocalTime           `gorm:"-" json:"create_at"`           // comment 创建时间
    UpdateAt            local.LocalTime           `gorm:"-" json:"update_at"`           // comment 更新时间

}

func (*Config) TableName() string {
	return TableNameConfig
}
