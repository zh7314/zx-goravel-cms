package models

import "github.com/goravel/framework/support/carbon"

const TableNameMessage = "message"

type Message struct {
	ID       int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	Content  string           `gorm:"column:content" json:"content"`                     // comment 内容
	CreateAt *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Email    string           `gorm:"column:email" json:"email"`                         // comment 电子邮箱
	Ip       string           `gorm:"column:ip" json:"ip"`                               // comment ip
	IsSent   int              `gorm:"column:is_sent" json:"is_sent"`                     // comment 10默认100成功
	Lang     string           `gorm:"column:lang" json:"lang"`                           // comment 语言类型
	Mobile   string           `gorm:"column:mobile" json:"mobile"`                       // comment 手机
	Pics     string           `gorm:"column:pics" json:"pics"`                           // comment 上传图片路径
	Platform string           `gorm:"column:platform" json:"platform"`                   // comment 平台类型
	RealName string           `gorm:"column:real_name" json:"real_name"`                 // comment 联系人名称
	Remark   string           `gorm:"column:remark" json:"remark"`                       // comment 处理备注
	Status   int              `gorm:"column:status" json:"status"`                       // comment 10未处理20已处理
	Title    string           `gorm:"column:title" json:"title"`                         // comment 标题
	Type     int              `gorm:"column:type" json:"type"`                           // comment 10免费找货，20集中采购，30委托寄售
	UpdateAt *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间

}

func (*Message) TableName() string {
	return TableNameMessage
}
