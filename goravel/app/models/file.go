package models

import "github.com/goravel/framework/support/carbon"

const TableNameFile = "file"

type File struct {
	ID       int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	AdminId  int64            `gorm:"column:admin_id" json:"admin_id"`                   // comment 创建者ID
	CreateAt *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	FileName string           `gorm:"column:file_name" json:"file_name"`                 // comment 文件名称
	FilePath string           `gorm:"column:file_path" json:"file_path"`                 // comment 文件路径
	FileSize int              `gorm:"column:file_size" json:"file_size"`                 // comment 文件大小

}

func (*File) TableName() string {
	return TableNameFile
}
