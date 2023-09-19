package admin

type ConfigRequest struct {
	ID       int64  `form:"id" json:"id"`             // ID
	Page     int    `form:"page" json:"page"`         // 页码
	PageSize int    `form:"pageSize" json:"pageSize"` // 页大小
	Name     string `form:"name" json:"name"`         // comment 名称
	Type     string `form:"type" json:"type"`         // comment 标签/类型
	Value    string `form:"value" json:"value"`       // comment 值

}
