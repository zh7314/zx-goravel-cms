package models

import "github.com/goravel/framework/support/carbon"

const TableNameVideo = "video"

type Video struct {
	ID          int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	AdminId     int              `gorm:"column:admin_id" json:"admin_id"`                   // comment 创建者ID
	Count       int              `gorm:"column:count" json:"count"`                         // comment 浏览数
	CreateAt    *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	File        string           `gorm:"column:file" json:"file"`                           // comment 视频文件url
	Introduce   string           `gorm:"column:introduce" json:"introduce"`                 // comment 视频介绍
	IsLocal     int              `gorm:"column:is_local" json:"is_local"`                   // comment 上传到本地服务器10是，20是外部视频
	IsShow      int              `gorm:"column:is_show" json:"is_show"`                     // comment 是否显示10显示20不显示
	Lang        string           `gorm:"column:lang" json:"lang"`                           // comment 语言类型
	Name        string           `gorm:"column:name" json:"name"`                           // comment 视频名称
	Pic         string           `gorm:"column:pic" json:"pic"`                             // comment 封面图片url
	Platform    string           `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	Sort        int              `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	UpdateAt    *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Url         string           `gorm:"column:url" json:"url"`                             // comment 外部视频URL
	VideoCateId int64            `gorm:"column:video_cate_id" json:"video_cate_id"`         // comment 分类视频ID
	VideoCate   *VideoCate       `gorm:"foreignKey:video_cate_id;references:id" json:"video_cate"`
}

func (*Video) TableName() string {
	return TableNameVideo
}
