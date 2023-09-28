package models

import (
	"github.com/goravel/framework/support/carbon"
	"time"
)

const TableNameAdmin = "admin"

type Admin struct {
	ID            int64            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // comment ID
	AdminGroupIds string           `gorm:"column:admin_group_ids" json:"admin_group_ids"`     // comment 用户权限组ID集合
	Avatar        string           `gorm:"column:avatar" json:"avatar"`                       // comment 用户头像
	CreateAt      *carbon.DateTime `gorm:"column:create_at;->" json:"create_at"`              // comment 创建时间
	Email         string           `gorm:"column:email" json:"email"`                         // comment 邮件
	IsAdmin       int              `gorm:"column:is_admin" json:"is_admin"`                   // comment 是否管理员 10是，99不是
	LoginIp       string           `gorm:"column:login_ip" json:"login_ip"`                   // comment 登陆IP
	Mobile        string           `gorm:"column:mobile" json:"mobile"`                       // comment 手机号码
	Name          string           `gorm:"column:name" json:"name"`                           // comment 系统管理用户名
	Password      string           `gorm:"column:password" json:"password"`                   // comment 密码
	RealName      string           `gorm:"column:real_name" json:"real_name"`                 // comment 真是姓名
	Salt          string           `gorm:"column:salt" json:"salt"`                           // comment 加密盐
	Sex           int              `gorm:"column:sex" json:"sex"`                             // comment 10保密20男30女
	Sort          int              `gorm:"column:sort" json:"sort"`                           // comment 排序越小越往前
	Status        int              `gorm:"column:status" json:"status"`                       // comment 用户状态10是默认正常 20是禁止登陆
	Token         string           `gorm:"column:token" json:"token"`                         // comment token
	TokenTime     *time.Time       `gorm:"column:token_time" json:"token_time"`               // comment 登陆token时间
	UpdateAt      *carbon.DateTime `gorm:"column:update_at;->" json:"update_at"`              // comment 更新时间

}

func (*Admin) TableName() string {
	return TableNameAdmin
}
