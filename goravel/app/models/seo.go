package models

import "github.com/goravel/framework/support/carbon"

const TableNameSeo = "seo"

type Seo struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    Title            string           `gorm:"column:title" json:"title"`           // comment 标题
    Keyword            string           `gorm:"column:keyword" json:"keyword"`           // comment 关键词
    Description            string           `gorm:"column:description" json:"description"`           // comment 描述
    Position            string           `gorm:"column:position" json:"position"`           // comment 位置
    CreateAt            carbon.DateTime           `gorm:"column:create_at;->" json:"create_at"`           // comment 创建时间
    UpdateAt            carbon.DateTime           `gorm:"column:update_at;->" json:"update_at"`           // comment 更新时间
    IsShow            int           `gorm:"column:is_show" json:"is_show"`           // comment 是否显示10是20不显示	
    Sort            int           `gorm:"column:sort" json:"sort"`           // comment 排序越小越往前
    Platform            string           `gorm:"column:platform" json:"platform"`           // comment 平台类型
    Lang            string           `gorm:"column:lang" json:"lang"`           // comment 语言类型

}

func (*Seo) TableName() string {
	return TableNameSeo
}