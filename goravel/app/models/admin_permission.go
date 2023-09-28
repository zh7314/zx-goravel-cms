package models

import "github.com/goravel/framework/support/carbon"

const TableNameAdminPermission = "admin_permission"

type AdminPermission struct {
	ID        int64              `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	Component string             `gorm:"column:component" json:"component"`                 // comment vue前台页面
	CreateAt  *carbon.DateTime   `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Hidden    int                `gorm:"column:hidden" json:"hidden"`                       // comment 10显示20不显示
	Icon      string             `gorm:"column:icon" json:"icon"`                           // comment 小图标名称
	IsMenu    int                `gorm:"column:is_menu" json:"is_menu"`                     // comment 作为菜单显示,10是,20不是
	Name      string             `gorm:"column:name" json:"name"`                           // comment 控制名称
	ParentId  int64              `gorm:"column:parent_id" json:"parent_id"`                 // comment 父ID 0是顶级
	Path      string             `gorm:"column:path" json:"path"`                           // comment 控制器URL
	Sort      int                `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	UpdateAt  *carbon.DateTime   `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Children  []*AdminPermission `gorm:"foreignKey:parent_id" json:"children"`
}

func (*AdminPermission) TableName() string {
	return TableNameAdminPermission
}
