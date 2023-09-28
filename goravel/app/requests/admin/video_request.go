package admin

type VideoRequest struct {
	ID          int64  `form:"id" json:"id"`                       // ID
	Page        int    `form:"page" json:"page"`                   // 页码
	PageSize    int    `form:"pageSize" json:"pageSize"`           // 页大小
	AdminId     int    `form:"admin_id" json:"admin_id"`           // comment 创建者ID
	Count       int    `form:"count" json:"count"`                 // comment 浏览数
	File        string `form:"file" json:"file"`                   // comment 视频文件url
	Introduce   string `form:"introduce" json:"introduce"`         // comment 视频介绍
	IsLocal     int    `form:"is_local" json:"is_local"`           // comment 上传到本地服务器10是，20是外部视频
	IsShow      int    `form:"is_show" json:"is_show"`             // comment 是否显示10显示20不显示
	Lang        string `form:"lang" json:"lang"`                   // comment 语言类型
	Name        string `form:"name" json:"name"`                   // comment 视频名称
	Pic         string `form:"pic" json:"pic"`                     // comment 封面图片url
	Platform    string `form:"platform" json:"platform"`           // comment 平台类型
	Sort        int    `form:"sort" json:"sort"`                   // comment 排序越小越往前
	Url         string `form:"url" json:"url"`                     // comment 外部视频URL
	VideoCateId int64  `form:"video_cate_id" json:"video_cate_id"` // comment 分类视频ID

}
