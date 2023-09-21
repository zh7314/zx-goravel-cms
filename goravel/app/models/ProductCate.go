package models

import "goravel/app/utils/local"

const TableNameProductCate = "product_cate"

type ProductCate struct {
	ID          int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	ParentId    int64           `gorm:"column:parent_id" json:"parent_id"`                 // comment 父产品分类ID
	Name        string          `gorm:"column:name" json:"name"`                           // comment 产品分类名称
	IsShow      int             `gorm:"column:is_show" json:"is_show"`                     // comment 是否显示10是显示，20不是
	Sort        int             `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	Url         string          `gorm:"column:url" json:"url"`                             // comment 产品分类外部url
	CreateAt    local.LocalTime `gorm:"-" json:"create_at"`                                // comment 创建时间
	UpdateAt    local.LocalTime `gorm:"-" json:"update_at"`                                // comment 更新时间
	Description string          `gorm:"column:description" json:"description"`             // comment 分类描述
	Pic         string          `gorm:"column:pic" json:"pic"`                             // comment 产品分类图片url
	Platform    string          `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	Lang        string          `gorm:"column:lang" json:"lang"`                           // comment 语言类型

}

func (*ProductCate) TableName() string {
	return TableNameProductCate
}
