package admin

type VideoCateRequest struct {
	ID       int64  `form:"id" json:"id"`               // ID
	Page     int    `form:"page" json:"page"`           // 页码
	PageSize int    `form:"pageSize" json:"pageSize"`   // 页大小
	ParentId int64  `form:"parent_id" json:"parent_id"` // comment 父分类名称
	Name     string `form:"name" json:"name"`           // comment 视频分类名称
	IsShow   int    `form:"is_show" json:"is_show"`     // comment 是否显示10显示20不显示
	Sort     int    `form:"sort" json:"sort"`           // comment 排序越小越往前
	Platform string `form:"platform" json:"platform"`   // comment 平台类型
	Lang     string `form:"lang" json:"lang"`           // comment 语言类型

}
