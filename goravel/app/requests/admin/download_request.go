package admin



type DownloadRequest struct {
    ID            int64           `form:"id" json:"id"`                           // ID
	Page          int             `form:"page" json:"page"`                       // 页码
	PageSize      int             `form:"pageSize" json:"pageSize"`               // 页大小
    Name            string           `form:"name" json:"name"`           // comment 文件名称
    Introduction            string           `form:"introduction" json:"introduction"`           // comment 文件简介
    Url            string           `form:"url" json:"url"`           // comment 连接地址
    IsLocal            int           `form:"is_local" json:"is_local"`           // comment 上传到本地服务器10是，20是外部地址	
    AdminId            int64           `form:"admin_id" json:"admin_id"`           // comment 管理员ID
    IsShow            int           `form:"is_show" json:"is_show"`           // comment 是否显示10是20不显示
    Sort            int           `form:"sort" json:"sort"`           // comment 排序越小越往前
    Path            string           `form:"path" json:"path"`           // comment 文件路径url
    DownloadCateId            int64           `form:"download_cate_id" json:"download_cate_id"`           // comment 分类id
    Pic            string           `form:"pic" json:"pic"`           // comment 封面图url
    Platform            string           `form:"platform" json:"platform"`           // comment 平台类型
    Lang            string           `form:"lang" json:"lang"`           // comment 语言类型

}
