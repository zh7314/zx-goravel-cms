package models

import "goravel/app/utils/local"

const TableNameNews = "news"

type News struct {
	ID             int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	NewsCateId     int64           `gorm:"column:news_cate_id" json:"news_cate_id"`           // comment 新闻分类ID
	Title          string          `gorm:"column:title" json:"title"`                         // comment 新闻标题
	ShortTitle     string          `gorm:"column:short_title" json:"short_title"`             // comment 新闻短标题
	Content        string          `gorm:"column:content" json:"content"`                     // comment 新闻内容
	AdminId        int             `gorm:"column:admin_id" json:"admin_id"`                   // comment 创建者ID
	Count          int             `gorm:"column:count" json:"count"`                         // comment 浏览数
	IsShow         int             `gorm:"column:is_show" json:"is_show"`                     // comment 10是默认显示，20是不显示
	Sort           int             `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	StartTime      local.LocalTime `gorm:"column:start_time" json:"start_time"`               // comment 开始时间
	EndTime        local.LocalTime `gorm:"column:end_time" json:"end_time"`                   // comment 结束时间
	CreateAt       local.LocalTime `gorm:"-" json:"create_at"`                                // comment 创建时间
	UpdateAt       local.LocalTime `gorm:"-" json:"update_at"`                                // comment 更新时间
	Pic            string          `gorm:"column:pic" json:"pic"`                             // comment 图片大图
	Pics           string          `gorm:"column:pics" json:"pics"`                           // comment 新闻图集
	SeoTitle       string          `gorm:"column:seo_title" json:"seo_title"`                 // comment seo标题
	SeoKeyword     string          `gorm:"column:seo_keyword" json:"seo_keyword"`             // comment seo关键词
	SeoDescription string          `gorm:"column:seo_description" json:"seo_description"`     // comment seo描述
	Platform       string          `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	Lang           string          `gorm:"column:lang" json:"lang"`                           // comment 语言类型

}

func (*News) TableName() string {
	return TableNameNews
}
