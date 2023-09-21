package models

import "goravel/app/utils/local"

const TableNameAdminLog = "admin_log"

type AdminLog struct {
	ID        int64           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	CreateAt  local.LocalTime `gorm:"-" json:"create_at"`                                // comment 创建时间
	UpdateAt  local.LocalTime `gorm:"-" json:"update_at"`                                // comment 更新时间
	Method    string          `gorm:"column:method" json:"method"`                       // comment 请求方式
	Url       string          `gorm:"column:url" json:"url"`                             // comment 请求url带参数
	RouteName string          `gorm:"column:route_name" json:"route_name"`               // comment 框架里定义的路由名称
	Path      string          `gorm:"column:path" json:"path"`                           // comment 请求url不带参数
	RequestIp string          `gorm:"column:request_ip" json:"request_ip"`               // comment 请求ip
	Data      string          `gorm:"column:data" json:"data"`                           // comment 请求参数
	AdminId   int64           `gorm:"column:admin_id" json:"admin_id"`                   // comment 管理员ID
	AdminName string          `gorm:"column:admin_name" json:"admin_name"`               // comment 管理员名称
	RouteDesc string          `gorm:"column:route_desc" json:"route_desc"`               // comment 路由描述

}

func (*AdminLog) TableName() string {
	return TableNameAdminLog
}
