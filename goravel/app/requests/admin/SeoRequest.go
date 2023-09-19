package admin

type SeoRequest struct {
	ID          int64  `form:"id" json:"id"`                   // ID
	Page        int    `form:"page" json:"page"`               // 页码
	PageSize    int    `form:"pageSize" json:"pageSize"`       // 页大小
	Description string `form:"description" json:"description"` // comment 描述
	IsShow      int    `form:"is_show" json:"is_show"`         // comment 是否显示10是20不显示
	Keyword     string `form:"keyword" json:"keyword"`         // comment 关键词
	Lang        string `form:"lang" json:"lang"`               // comment 语言类型
	Platform    string `form:"platform" json:"platform"`       // comment 平台类型
	Position    string `form:"position" json:"position"`       // comment 位置
	Sort        int    `form:"sort" json:"sort"`               // comment 排序越小越往前
	Title       string `form:"title" json:"title"`             // comment 标题

}
