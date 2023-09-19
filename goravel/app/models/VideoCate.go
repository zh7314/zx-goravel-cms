package models

import (
	"goravel/app/utils/local"
)

const TableNameVideoCate = "video_cate"

type VideoCate struct {
	ID       int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	CreateAt local.LocalTime `gorm:"-" json:"create_at"`                                // comment 创建时间
	IsShow   int             `gorm:"column:is_show" json:"is_show"`                     // comment 是否显示10显示20不显示
	Lang     string          `gorm:"column:lang" json:"lang"`                           // comment 语言类型
	Name     string          `gorm:"column:name" json:"name"`                           // comment 视频分类名称
	ParentId int64           `gorm:"column:parent_id" json:"parent_id"`                 // comment 父分类名称
	Platform string          `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	Sort     int             `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	UpdateAt local.LocalTime `gorm:"-" json:"update_at"`                                // comment 更新时间

}

func (*VideoCate) TableName() string {
	return TableNameVideoCate
}
