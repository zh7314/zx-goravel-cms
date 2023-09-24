package models

import "github.com/goravel/framework/support/carbon"

const TableNameVideo = "video"

type Video struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    VideoCateId            int64           `gorm:"column:video_cate_id" json:"video_cate_id"`           // comment 分类视频ID
    IsLocal            int           `gorm:"column:is_local" json:"is_local"`           // comment 上传到本地服务器10是，20是外部视频
    Url            string           `gorm:"column:url" json:"url"`           // comment 外部视频URL
    Name            string           `gorm:"column:name" json:"name"`           // comment 视频名称
    Introduce            string           `gorm:"column:introduce" json:"introduce"`           // comment 视频介绍
    IsShow            int           `gorm:"column:is_show" json:"is_show"`           // comment 是否显示10显示20不显示
    AdminId            int           `gorm:"column:admin_id" json:"admin_id"`           // comment 创建者ID
    Count            int           `gorm:"column:count" json:"count"`           // comment 浏览数
    Sort            int           `gorm:"column:sort" json:"sort"`           // comment 排序越小越往前
    Pic            string           `gorm:"column:pic" json:"pic"`           // comment 封面图片url
    File            string           `gorm:"column:file" json:"file"`           // comment 视频文件url
    CreateAt            carbon.DateTime           `gorm:"column:create_at;->" json:"create_at"`           // comment 创建时间
    UpdateAt            carbon.DateTime           `gorm:"column:update_at;->" json:"update_at"`           // comment 更新时间
    Platform            string           `gorm:"column:platform" json:"platform"`           // comment 平台类型
    Lang            string           `gorm:"column:lang" json:"lang"`           // comment 语言类型

}

func (*Video) TableName() string {
	return TableNameVideo
}
