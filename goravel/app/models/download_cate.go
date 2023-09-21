package models

import "goravel/app/utils/local"

const TableNameDownloadCate = "download_cate"

type DownloadCate struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    Name            string           `gorm:"column:name" json:"name"`           // comment 分类名称
    IsShow            int           `gorm:"column:is_show" json:"is_show"`           // comment 是否显示10显示20不显示	
    Sort            int           `gorm:"column:sort" json:"sort"`           // comment 排序越小越往前
    CreateAt            local.LocalTime           `gorm:"-" json:"create_at"`           // comment 创建时间
    UpdateAt            local.LocalTime           `gorm:"-" json:"update_at"`           // comment 更新时间
    Pic            string           `gorm:"column:pic" json:"pic"`           // comment 分类图片地址
    Platform            string           `gorm:"column:platform" json:"platform"`           // comment 平台类型
    Lang            string           `gorm:"column:lang" json:"lang"`           // comment 语言类型
    ParentId            int64           `gorm:"column:parent_id" json:"parent_id"`           // comment 父级id

}

func (*DownloadCate) TableName() string {
	return TableNameDownloadCate
}
