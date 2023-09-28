package admin

type AdminGroupRequest struct {
	ID            int64  `form:"id" json:"id"`                         // ID
	Page          int    `form:"page" json:"page"`                     // 页码
	PageSize      int    `form:"pageSize" json:"pageSize"`             // 页大小
	Name          string `form:"name" json:"name"`                     // comment 分组名称
	ParentId      int64  `form:"parent_id" json:"parent_id"`           // comment 父ID 0是顶级
	PermissionIds []int  `form:"permission_ids" json:"permission_ids"` // comment permission_id集合
	Sort          int    `form:"sort" json:"sort"`                     // comment 排序越小越往前

}
