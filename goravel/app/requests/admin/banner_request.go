package admin

import "time"

type BannerRequest struct {
	ID           int64     `form:"id" json:"id"`                         // ID
	Page         int       `form:"page" json:"page"`                     // 页码
	PageSize     int       `form:"pageSize" json:"pageSize"`             // 页大小
	AdminId      int64     `form:"admin_id" json:"admin_id"`             // comment 创建ID
	BannerCateId int64     `form:"banner_cate_id" json:"banner_cate_id"` // comment 分类id
	EndTime      time.Time `form:"end_time" json:"end_time"`             // comment 结束时间
	Lang         string    `form:"lang" json:"lang"`                     // comment 语言类型默认
	Name         string    `form:"name" json:"name"`                     // comment banner名称
	PicPath      string    `form:"pic_path" json:"pic_path"`             // comment 图片地址
	Platform     string    `form:"platform" json:"platform"`             // comment 平台类型
	Sort         int       `form:"sort" json:"sort"`                     // comment 排序越小越往前
	StartTime    time.Time `form:"start_time" json:"start_time"`         // comment 开始时间
	Url          string    `form:"url" json:"url"`                       // comment 跳转链接
	VideoPath    string    `form:"video_path" json:"video_path"`         // comment 视频地址

}
