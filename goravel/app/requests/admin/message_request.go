package admin



type MessageRequest struct {
    ID            int64           `form:"id" json:"id"`                           // ID
	Page          int             `form:"page" json:"page"`                       // 页码
	PageSize      int             `form:"pageSize" json:"pageSize"`               // 页大小
    Type            int           `form:"type" json:"type"`           // comment 10免费找货，20集中采购，30委托寄售
    Mobile            string           `form:"mobile" json:"mobile"`           // comment 手机
    RealName            string           `form:"real_name" json:"real_name"`           // comment 联系人名称
    Email            string           `form:"email" json:"email"`           // comment 电子邮箱
    Ip            string           `form:"ip" json:"ip"`           // comment ip
    Status            int           `form:"status" json:"status"`           // comment 10未处理20已处理
    Title            string           `form:"title" json:"title"`           // comment 标题
    Content            string           `form:"content" json:"content"`           // comment 内容
    Pics            string           `form:"pics" json:"pics"`           // comment 上传图片路径
    IsSent            int           `form:"is_sent" json:"is_sent"`           // comment 10默认100成功
    Remark            string           `form:"remark" json:"remark"`           // comment 处理备注
    Platform            string           `form:"platform" json:"platform"`           // comment 平台类型
    Lang            string           `form:"lang" json:"lang"`           // comment 语言类型

}
