package admin

type ProductCateRequest struct {
	ID          int64  `form:"id" json:"id"`                   // ID
	Page        int    `form:"page" json:"page"`               // 页码
	PageSize    int    `form:"pageSize" json:"pageSize"`       // 页大小
	Description string `form:"description" json:"description"` // comment 分类描述
	IsShow      int    `form:"is_show" json:"is_show"`         // comment 是否显示10是显示，20不是
	Lang        string `form:"lang" json:"lang"`               // comment 语言类型
	Name        string `form:"name" json:"name"`               // comment 产品分类名称
	ParentId    int64  `form:"parent_id" json:"parent_id"`     // comment 父产品分类ID
	Pic         string `form:"pic" json:"pic"`                 // comment 产品分类图片url
	Platform    string `form:"platform" json:"platform"`       // comment 平台类型
	Sort        int    `form:"sort" json:"sort"`               // comment 排序越小越往前
	Url         string `form:"url" json:"url"`                 // comment 产品分类外部url

}
