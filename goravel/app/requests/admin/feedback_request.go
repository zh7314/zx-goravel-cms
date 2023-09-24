package admin



type FeedbackRequest struct {
    ID            int64           `form:"id" json:"id"`                           // ID
	Page          int             `form:"page" json:"page"`                       // 页码
	PageSize      int             `form:"pageSize" json:"pageSize"`               // 页大小
    NickName            string           `form:"nick_name" json:"nick_name"`           // comment 反馈昵称
    Contact            string           `form:"contact" json:"contact"`           // comment 联系方式
    Content            string           `form:"content" json:"content"`           // comment 反馈内容
    Platform            string           `form:"platform" json:"platform"`           // comment 平台类型
    Lang            string           `form:"lang" json:"lang"`           // comment 语言类型

}
