package models

import (
	"goravel/app/utils/local"
)

const TableNameBanner = "banner"

type Banner struct {
	ID           int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	AdminId      int64           `gorm:"column:admin_id" json:"admin_id"`                   // comment 创建ID
	BannerCateId int64           `gorm:"column:banner_cate_id" json:"banner_cate_id"`       // comment 分类id
	CreateAt     local.LocalTime `gorm:"-" json:"create_at"`                                // comment 创建时间
	EndTime      local.LocalTime `gorm:"column:end_time" json:"end_time"`                   // comment 结束时间
	Lang         string          `gorm:"column:lang" json:"lang"`                           // comment 语言类型默认
	Name         string          `gorm:"column:name" json:"name"`                           // comment banner名称
	PicPath      string          `gorm:"column:pic_path" json:"pic_path"`                   // comment 图片地址
	Platform     string          `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	Sort         int             `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	StartTime    local.LocalTime `gorm:"column:start_time" json:"start_time"`               // comment 开始时间
	UpdateAt     local.LocalTime `gorm:"-" json:"update_at"`                                // comment 更新时间
	Url          string          `gorm:"column:url" json:"url"`                             // comment 跳转链接
	VideoPath    string          `gorm:"column:video_path" json:"video_path"`               // comment 视频地址

}

func (*Banner) TableName() string {
	return TableNameBanner
}
