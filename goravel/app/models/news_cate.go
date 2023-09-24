package models

import "github.com/goravel/framework/support/carbon"

const TableNameNewsCate = "news_cate"

type NewsCate struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    ParentId            int64           `gorm:"column:parent_id" json:"parent_id"`           // comment 父分类ID
    Name            string           `gorm:"column:name" json:"name"`           // comment 新闻分类名称
    IsShow            int           `gorm:"column:is_show" json:"is_show"`           // comment 10是默认显示，20是不显示
    Sort            int           `gorm:"column:sort" json:"sort"`           // comment 排序越小越往前
    CreateAt            carbon.DateTime           `gorm:"column:create_at;->" json:"create_at"`           // comment 创建时间
    UpdateAt            carbon.DateTime           `gorm:"column:update_at;->" json:"update_at"`           // comment 更新时间
    Platform            string           `gorm:"column:platform" json:"platform"`           // comment 平台类型
    Lang            string           `gorm:"column:lang" json:"lang"`           // comment 语言类型

}

func (*NewsCate) TableName() string {
	return TableNameNewsCate
}
