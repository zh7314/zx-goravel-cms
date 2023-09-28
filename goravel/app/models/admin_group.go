package models

import "github.com/goravel/framework/support/carbon"

const TableNameAdminGroup = "admin_group"

type AdminGroup struct {
	ID            int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	CreateAt      *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Name          string           `gorm:"column:name" json:"name"`                           // comment 分组名称
	ParentId      int64            `gorm:"column:parent_id" json:"parent_id"`                 // comment 父ID 0是顶级
	PermissionIds string           `gorm:"column:permission_ids" json:"permission_ids"`       // comment permission_id集合
	Sort          int              `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	UpdateAt      *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Children      []*AdminGroup    `gorm:"foreignKey:parent_id" json:"children"`
}

func (*AdminGroup) TableName() string {
	return TableNameAdminGroup
}
