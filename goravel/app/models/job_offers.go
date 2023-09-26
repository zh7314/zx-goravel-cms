package models

import "github.com/goravel/framework/support/carbon"

const TableNameJobOffers = "job_offers"

type JobOffers struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
    Title            string           `gorm:"column:title" json:"title"`           // comment 招聘工种
    Url            string           `gorm:"column:url" json:"url"`           // comment 招聘链接URL
    Content            string           `gorm:"column:content" json:"content"`           // comment 招聘内容
    IsShow            int           `gorm:"column:is_show" json:"is_show"`           // comment 10是显示，20是不显示
    Sort            int           `gorm:"column:sort" json:"sort"`           // comment 排序
    SalaryRange            string           `gorm:"column:salary_range" json:"salary_range"`           // comment 薪资范围
    Place            string           `gorm:"column:place" json:"place"`           // comment 工作地点
    Number            string           `gorm:"column:number" json:"number"`           // comment 招聘人数
    CreateAt            carbon.DateTime           `gorm:"column:create_at;->" json:"create_at"`           // comment 创建时间
    UpdateAt            carbon.DateTime           `gorm:"column:update_at;->" json:"update_at"`           // comment 更新时间
    Platform            string           `gorm:"column:platform" json:"platform"`           // comment 平台类型
    Lang            string           `gorm:"column:lang" json:"lang"`           // comment 语言类型

}

func (*JobOffers) TableName() string {
	return TableNameJobOffers
}
