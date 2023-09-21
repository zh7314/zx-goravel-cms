package admin

type NewsRequest struct {
	ID             int64  `form:"id" json:"id"`                           // ID
	Page           int    `form:"page" json:"page"`                       // 页码
	PageSize       int    `form:"pageSize" json:"pageSize"`               // 页大小
	NewsCateId     int64  `form:"news_cate_id" json:"news_cate_id"`       // comment 新闻分类ID
	Title          string `form:"title" json:"title"`                     // comment 新闻标题
	ShortTitle     string `form:"short_title" json:"short_title"`         // comment 新闻短标题
	Content        string `form:"content" json:"content"`                 // comment 新闻内容
	AdminId        int    `form:"admin_id" json:"admin_id"`               // comment 创建者ID
	Count          int    `form:"count" json:"count"`                     // comment 浏览数
	IsShow         int    `form:"is_show" json:"is_show"`                 // comment 10是默认显示，20是不显示
	Sort           int    `form:"sort" json:"sort"`                       // comment 排序越小越往前
	StartTime      string `form:"start_time" json:"start_time"`           // comment 开始时间
	EndTime        string `form:"end_time" json:"end_time"`               // comment 结束时间
	Pic            string `form:"pic" json:"pic"`                         // comment 图片大图
	Pics           string `form:"pics" json:"pics"`                       // comment 新闻图集
	SeoTitle       string `form:"seo_title" json:"seo_title"`             // comment seo标题
	SeoKeyword     string `form:"seo_keyword" json:"seo_keyword"`         // comment seo关键词
	SeoDescription string `form:"seo_description" json:"seo_description"` // comment seo描述
	Platform       string `form:"platform" json:"platform"`               // comment 平台类型
	Lang           string `form:"lang" json:"lang"`                       // comment 语言类型
}
