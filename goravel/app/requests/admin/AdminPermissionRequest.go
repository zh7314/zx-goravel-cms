package admin

type AdminPermissionRequest struct {
	ID        int64  `form:"id" json:"id"`               // ID
	Page      int    `form:"page" json:"page"`           // 页码
	PageSize  int    `form:"pageSize" json:"pageSize"`   // 页大小
	Component string `form:"component" json:"component"` // comment vue前台页面
	Hidden    int    `form:"hidden" json:"hidden"`       // comment 10显示20不显示
	Icon      string `form:"icon" json:"icon"`           // comment 小图标名称
	IsMenu    int    `form:"is_menu" json:"is_menu"`     // comment 作为菜单显示,10是,20不是
	Name      string `form:"name" json:"name"`           // comment 控制名称
	ParentId  int64  `form:"parent_id" json:"parent_id"` // comment 父ID 0是顶级
	Path      string `form:"path" json:"path"`           // comment 控制器URL
	Sort      int    `form:"sort" json:"sort"`           // comment 排序越小越往前

}
