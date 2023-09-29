package models

import "github.com/goravel/framework/support/carbon"

const TableNameAdminLog = "admin_log"

type AdminLog struct {
	ID        int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	AdminId   int64            `gorm:"column:admin_id" json:"admin_id"`                   // comment 管理员ID
	AdminName string           `gorm:"column:admin_name" json:"admin_name"`               // comment 管理员名称
	CreateAt  *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Data      string           `gorm:"column:data" json:"data"`                           // comment 请求参数
	Method    string           `gorm:"column:method" json:"method"`                       // comment 请求方式
	Path      string           `gorm:"column:path" json:"path"`                           // comment 请求url不带参数
	RequestIp string           `gorm:"column:request_ip" json:"request_ip"`               // comment 请求ip
	RouteDesc string           `gorm:"column:route_desc" json:"route_desc"`               // comment 路由描述
	RouteName string           `gorm:"column:route_name" json:"route_name"`               // comment 框架里定义的路由名称
	UpdateAt  *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间
	Url       string           `gorm:"column:url" json:"url"`                             // comment 请求url带参数
}

func (*AdminLog) TableName() string {
	return TableNameAdminLog
}
