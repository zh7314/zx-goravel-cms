package admin



type JobOffersRequest struct {
    ID            int64           `form:"id" json:"id"`                           // ID
	Page          int             `form:"page" json:"page"`                       // 页码
	PageSize      int             `form:"pageSize" json:"pageSize"`               // 页大小
    Title            string           `form:"title" json:"title"`           // comment 招聘工种
    Url            string           `form:"url" json:"url"`           // comment 招聘链接URL
    Content            string           `form:"content" json:"content"`           // comment 招聘内容
    IsShow            int           `form:"is_show" json:"is_show"`           // comment 10是显示，20是不显示
    Sort            int           `form:"sort" json:"sort"`           // comment 排序
    SalaryRange            string           `form:"salary_range" json:"salary_range"`           // comment 薪资范围
    Place            string           `form:"place" json:"place"`           // comment 工作地点
    Number            string           `form:"number" json:"number"`           // comment 招聘人数
    Platform            string           `form:"platform" json:"platform"`           // comment 平台类型
    Lang            string           `form:"lang" json:"lang"`           // comment 语言类型

}
