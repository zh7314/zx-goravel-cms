package admin

type FileRequest struct {
	ID       int64  `form:"id" json:"id"`               // ID
	Page     int    `form:"page" json:"page"`           // 页码
	PageSize int    `form:"pageSize" json:"pageSize"`   // 页大小
	AdminId  int64  `form:"admin_id" json:"admin_id"`   // comment 创建者ID
	FileName string `form:"file_name" json:"file_name"` // comment 文件名称
	FilePath string `form:"file_path" json:"file_path"` // comment 文件路径
	FileSize int    `form:"file_size" json:"file_size"` // comment 文件大小

}
