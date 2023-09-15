package models

import (
	"goravel/app/utils/local"
)

const TableNameBanner = "banner"

// Banner mapped from table <banner>
type Banner struct {
	ID           int64           `gorm:"column:id;primaryKey;autoIncrement:true;comment:bannerID" json:"id"`                // bannerID
	BannerCateID int64           `gorm:"column:banner_cate_id;not null;comment:分类id" json:"banner_cate_id"`                 // 分类id
	Name         string          `gorm:"column:name;not null;comment:banner名称" json:"name"`                                 // banner名称
	AdminID      int64           `gorm:"column:admin_id;comment:创建ID" json:"admin_id"`                                      // 创建ID
	URL          string          `gorm:"column:url;not null;comment:跳转链接" json:"url"`                                       // 跳转链接
	Sort         bool            `gorm:"column:sort;not null;default:255;comment:排序越小越往前" json:"sort"`                      // 排序越小越往前
	StartTime    local.LocalTime `gorm:"column:start_time;comment:开始时间" json:"start_time"`                                  // 开始时间
	EndTime      local.LocalTime `gorm:"column:end_time;comment:结束时间" json:"end_time"`                                      // 结束时间
	CreateAt     local.LocalTime `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"` // 创建时间
	UpdateAt     local.LocalTime `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_at"` // 更新时间
	PicPath      string          `gorm:"column:pic_path;comment:图片地址" json:"pic_path"`                                      // 图片地址
	VideoPath    string          `gorm:"column:video_path;comment:视频地址" json:"video_path"`                                  // 视频地址
	Platform     string          `gorm:"column:platform;default:PC;comment:平台类型" json:"platform"`                           // 平台类型
	Lang         string          `gorm:"column:lang;default:zh-CN;comment:语言类型默认" json:"lang"`                              // 语言类型默认
}

// TableName Banner's table name
func (*Banner) TableName() string {
	return TableNameBanner
}
