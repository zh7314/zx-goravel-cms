package models

import "goravel/app/utils/local"

const TableNamePlatform = "platform"

type Platform struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    Sort            int           `gorm:"column:sort" json:"sort"`           // comment 排序越小越往前
    CreateAt            local.LocalTime           `gorm:"-" json:"create_at"`           // comment 创建时间
    UpdateAt            local.LocalTime           `gorm:"-" json:"update_at"`           // comment 更新时间
    Name            string           `gorm:"column:name" json:"name"`           // comment 名称
    Value            string           `gorm:"column:value" json:"value"`           // comment 值

}

func (*Platform) TableName() string {
	return TableNamePlatform
}
