package models

import "goravel/app/utils/local"

const TableNameDownload = "download"

type Download struct {
	ID             int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	Name           string          `gorm:"column:name" json:"name"`                           // comment 文件名称
	Introduction   string          `gorm:"column:introduction" json:"introduction"`           // comment 文件简介
	Url            string          `gorm:"column:url" json:"url"`                             // comment 连接地址
	IsLocal        int             `gorm:"column:is_local" json:"is_local"`                   // comment 上传到本地服务器10是，20是外部地址
	AdminId        int64           `gorm:"column:admin_id" json:"admin_id"`                   // comment 管理员ID
	CreateAt       local.LocalTime `gorm:"-" json:"create_at"`                                // comment 创建时间
	UpdateAt       local.LocalTime `gorm:"-" json:"update_at"`                                // comment 更新时间
	IsShow         int             `gorm:"column:is_show" json:"is_show"`                     // comment 是否显示10是20不显示
	Sort           int             `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	Path           string          `gorm:"column:path" json:"path"`                           // comment 文件路径url
	DownloadCateId int64           `gorm:"column:download_cate_id" json:"download_cate_id"`   // comment 分类id
	Pic            string          `gorm:"column:pic" json:"pic"`                             // comment 封面图url
	Platform       string          `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	Lang           string          `gorm:"column:lang" json:"lang"`                           // comment 语言类型

}

func (*Download) TableName() string {
	return TableNameDownload
}
