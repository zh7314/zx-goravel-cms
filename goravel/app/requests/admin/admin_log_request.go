package admin

type AdminLogRequest struct {
	ID        int64  `form:"id" json:"id"`                 // ID
	Page      int    `form:"page" json:"page"`             // 页码
	PageSize  int    `form:"pageSize" json:"pageSize"`     // 页大小
	AdminId   int64  `form:"admin_id" json:"admin_id"`     // comment 管理员ID
	AdminName string `form:"admin_name" json:"admin_name"` // comment 管理员名称
	Data      string `form:"data" json:"data"`             // comment 请求参数
	Method    string `form:"method" json:"method"`         // comment 请求方式
	Path      string `form:"path" json:"path"`             // comment 请求url不带参数
	RequestIp string `form:"request_ip" json:"request_ip"` // comment 请求ip
	RouteDesc string `form:"route_desc" json:"route_desc"` // comment 路由描述
	RouteName string `form:"route_name" json:"route_name"` // comment 框架里定义的路由名称
	Url       string `form:"url" json:"url"`               // comment 请求url带参数

}
