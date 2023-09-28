package models

import (
	"github.com/goravel/framework/support/carbon"
	"time"
)

const TableNameProduct = "product"

type Product struct {
	ID            int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	AdminId       int              `gorm:"column:admin_id" json:"admin_id"`                   // comment 创建者ID
	Content       string           `gorm:"column:content" json:"content"`                     // comment 产品内容
	CreateAt      *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	EndTime       *time.Time       `gorm:"column:end_time" json:"end_time"`                   // comment 结束时间
	IsShow        int              `gorm:"column:is_show" json:"is_show"`                     // comment 10是默认显示，20是不显示
	Lang          string           `gorm:"column:lang" json:"lang"`                           // comment 语言类型
	Pic           string           `gorm:"column:pic" json:"pic"`                             // comment 产品首页大图文件
	Pics          string           `gorm:"column:pics" json:"pics"`                           // comment 产品多图文件集合
	Platform      string           `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	ProductCateId int              `gorm:"column:product_cate_id" json:"product_cate_id"`     // comment 产品分类ID
	ShortTitle    string           `gorm:"column:short_title" json:"short_title"`             // comment 产品短标题
	Sort          int              `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	StartTime     *time.Time       `gorm:"column:start_time" json:"start_time"`               // comment 开始时间
	Title         string           `gorm:"column:title" json:"title"`                         // comment 产品标题
	UpdateAt      *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Url           string           `gorm:"column:url" json:"url"`                             // comment 产品外部URL
	VideoUrl      string           `gorm:"column:video_url" json:"video_url"`                 // comment 视频地址
	ViewCount     int              `gorm:"column:view_count" json:"view_count"`               // comment 浏览数
	ProductCate   *ProductCate     `gorm:"foreignKey:product_cate_id;references:id" json:"product_cate"`
}

func (*Product) TableName() string {
	return TableNameProduct
}
