package admin

type RequestLogRequest struct {
	ID       int64  `form:"id" json:"id"`               // ID
	Page     int    `form:"page" json:"page"`           // 页码
	PageSize int    `form:"pageSize" json:"pageSize"`   // 页大小
	Method   string `form:"method" json:"method"`       // comment 请求方式
	Ip       string `form:"ip" json:"ip"`               // comment 请求ip
	Url      string `form:"url" json:"url"`             // comment 请求url
	Params   string `form:"params" json:"params"`       // comment 请求参数
	Header   string `form:"header" json:"header"`       // comment 请求header
	Data     string `form:"data" json:"data"`           // comment 返回的数据 不包含data
	ReturnAt string `form:"return_at" json:"return_at"` // comment 返回结果的时间

}
