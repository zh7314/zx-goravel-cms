package admin

type DownloadCateRequest struct {
	ID       int64  `form:"id" json:"id"`               // ID
	Page     int    `form:"page" json:"page"`           // 页码
	PageSize int    `form:"pageSize" json:"pageSize"`   // 页大小
	IsShow   int    `form:"is_show" json:"is_show"`     // comment 是否显示10显示20不显示
	Lang     string `form:"lang" json:"lang"`           // comment 语言类型
	Name     string `form:"name" json:"name"`           // comment 分类名称
	ParentId int64  `form:"parent_id" json:"parent_id"` // comment 父级id
	Pic      string `form:"pic" json:"pic"`             // comment 分类图片地址
	Platform string `form:"platform" json:"platform"`   // comment 平台类型
	Sort     int    `form:"sort" json:"sort"`           // comment 排序越小越往前

}
