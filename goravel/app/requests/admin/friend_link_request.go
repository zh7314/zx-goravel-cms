package admin

type FriendLinkRequest struct {
	ID       int64  `form:"id" json:"id"`               // ID
	Page     int    `form:"page" json:"page"`           // 页码
	PageSize int    `form:"pageSize" json:"pageSize"`   // 页大小
	IsFollow int    `form:"is_follow" json:"is_follow"` // comment 是否follow 10是，20不是
	IsShow   int    `form:"is_show" json:"is_show"`     // comment 10是显示，20是不显示
	Lang     string `form:"lang" json:"lang"`           // comment 语言类型
	Pic      string `form:"pic" json:"pic"`             // comment 图片地址
	Platform string `form:"platform" json:"platform"`   // comment 平台类型
	Sort     int    `form:"sort" json:"sort"`           // comment 排序越小越往前
	Title    string `form:"title" json:"title"`         // comment 标题
	Url      string `form:"url" json:"url"`             // comment URL

}
