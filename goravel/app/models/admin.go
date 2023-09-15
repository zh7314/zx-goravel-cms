package models

import (
	"goravel/app/utils/local"
)

const TableNameAdmin = "admin"

// Admin mapped from table <admin>
type Admin struct {
	ID            int64           `gorm:"column:id;primaryKey;autoIncrement:true;comment:用户ID" json:"id"`                          // 用户ID
	Name          string          `gorm:"column:name;comment:系统管理用户名" json:"name"`                                            // 系统管理用户名
	Password      string          `gorm:"column:password;comment:密码" json:"password"`                                              // 密码
	Salt          string          `gorm:"column:salt;comment:加密盐" json:"salt"`                                                    // 加密盐
	Sex           bool            `gorm:"column:sex;not null;default:10;comment:10保密20男30女" json:"sex"`                          // 10保密20男30女
	Email         string          `gorm:"column:email;comment:邮件" json:"email"`                                                    // 邮件
	Mobile        string          `gorm:"column:mobile;comment:手机号码" json:"mobile"`                                              // 手机号码
	LoginIP       string          `gorm:"column:login_ip;comment:登陆IP" json:"login_ip"`                                            // 登陆IP
	Status        bool            `gorm:"column:status;not null;default:10;comment:用户状态10是默认正常 20是禁止登陆" json:"status"` // 用户状态10是默认正常 20是禁止登陆
	Avatar        string          `gorm:"column:avatar;comment:用户头像" json:"avatar"`                                              // 用户头像
	RealName      string          `gorm:"column:real_name;comment:真是姓名" json:"real_name"`                                        // 真是姓名
	TokenTime     local.LocalTime `gorm:"column:token_time;comment:登陆token时间" json:"token_time"`                                 // 登陆token时间
	AdminGroupIds string          `gorm:"column:admin_group_ids;comment:用户权限组ID集合" json:"admin_group_ids"`                    // 用户权限组ID集合
	IsAdmin       bool            `gorm:"column:is_admin;not null;default:99;comment:是否管理员 10是，99不是" json:"is_admin"`        // 是否管理员 10是，99不是
	Sort          bool            `gorm:"column:sort;not null;default:255;comment:排序越小越往前" json:"sort"`                       // 排序越小越往前
	CreateAt      local.LocalTime `gorm:"column:create_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_at"`     // 创建时间
	UpdateAt      local.LocalTime `gorm:"column:update_at;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_at"`     // 更新时间
	Token         string          `gorm:"column:token;comment:token" json:"token"`                                                   // token
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
