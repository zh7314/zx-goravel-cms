package admin

import "goravel/app/utils/local"

type ProductRequest struct {
	ID            int64           `form:"id" json:"id"`                           // ID
	Page          int             `form:"page" json:"page"`                       // 页码
	PageSize      int             `form:"pageSize" json:"pageSize"`               // 页大小
	ProductCateId int             `form:"product_cate_id" json:"product_cate_id"` // comment 产品分类ID
	Title         string          `form:"title" json:"title"`                     // comment 产品标题
	ShortTitle    string          `form:"short_title" json:"short_title"`         // comment 产品短标题
	Content       string          `form:"content" json:"content"`                 // comment 产品内容
	AdminId       int             `form:"admin_id" json:"admin_id"`               // comment 创建者ID
	ViewCount     int             `form:"view_count" json:"view_count"`           // comment 浏览数
	IsShow        int             `form:"is_show" json:"is_show"`                 // comment 10是默认显示，20是不显示
	Sort          int             `form:"sort" json:"sort"`                       // comment 排序越小越往前
	Url           string          `form:"url" json:"url"`                         // comment 产品外部URL
	StartTime     local.LocalTime `form:"start_time" json:"start_time"`           // comment 开始时间
	EndTime       local.LocalTime `form:"end_time" json:"end_time"`               // comment 结束时间
	Pic           string          `form:"pic" json:"pic"`                         // comment 产品首页大图文件
	Pics          string          `form:"pics" json:"pics"`                       // comment 产品多图文件集合
	VideoUrl      string          `form:"video_url" json:"video_url"`             // comment 视频地址
	Platform      string          `form:"platform" json:"platform"`               // comment 平台类型
	Lang          string          `form:"lang" json:"lang"`                       // comment 语言类型

}
