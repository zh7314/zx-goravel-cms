package models

import "github.com/goravel/framework/support/carbon"

const TableNameFriendLink = "friend_link"

type FriendLink struct {
	ID       int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	CreateAt *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	IsFollow int              `gorm:"column:is_follow" json:"is_follow"`                 // comment 是否follow 10是，20不是
	IsShow   int              `gorm:"column:is_show" json:"is_show"`                     // comment 10是显示，20是不显示
	Lang     string           `gorm:"column:lang" json:"lang"`                           // comment 语言类型
	Pic      string           `gorm:"column:pic" json:"pic"`                             // comment 图片地址
	Platform string           `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	Sort     int              `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	Title    string           `gorm:"column:title" json:"title"`                         // comment 标题
	UpdateAt *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Url      string           `gorm:"column:url" json:"url"`                             // comment URL

}

func (*FriendLink) TableName() string {
	return TableNameFriendLink
}
