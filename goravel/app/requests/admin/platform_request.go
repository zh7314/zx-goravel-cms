package admin

type PlatformRequest struct {
	ID       int64  `form:"id" json:"id"`             // ID
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 页大小
	Name     string `form:"name" json:"name"`         // comment 名称
	Sort     int    `form:"sort" json:"sort"`         // comment 排序越小越往前
	Value    string `form:"value" json:"value"`       // comment 值

}
