package models

import "goravel/app/utils/local"

const TableNameLang = "lang"

type Lang struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    Sort            int           `gorm:"column:sort" json:"sort"`           // comment 排序越小越往前
    CreateAt            local.LocalTime           `gorm:"-" json:"create_at"`           // comment 创建时间
    UpdateAt            local.LocalTime           `gorm:"-" json:"update_at"`           // comment 更新时间
    Name            string           `gorm:"column:name" json:"name"`           // comment 语言名称
    Value            string           `gorm:"column:value" json:"value"`           // comment 语言值

}

func (*Lang) TableName() string {
	return TableNameLang
}
