package admin

import "time"

type AdminRequest struct {
	ID            int64     `form:"id" json:"id"`                           // ID
	Page          int       `form:"page" json:"page"`                       // 页码
	PageSize      int       `form:"pageSize" json:"pageSize"`               // 页大小
	AdminGroupIds []int     `form:"admin_group_ids" json:"admin_group_ids"` // comment 用户权限组ID集合
	Avatar        string    `form:"avatar" json:"avatar"`                   // comment 用户头像
	Email         string    `form:"email" json:"email"`                     // comment 邮件
	IsAdmin       int       `form:"is_admin" json:"is_admin"`               // comment 是否管理员 10是，99不是
	LoginIp       string    `form:"login_ip" json:"login_ip"`               // comment 登陆IP
	Mobile        string    `form:"mobile" json:"mobile"`                   // comment 手机号码
	Name          string    `form:"name" json:"name"`                       // comment 系统管理用户名
	Password      string    `form:"password" json:"password"`               // comment 密码
	RealName      string    `form:"real_name" json:"real_name"`             // comment 真是姓名
	Salt          string    `form:"salt" json:"salt"`                       // comment 加密盐
	Sex           int       `form:"sex" json:"sex"`                         // comment 10保密20男30女
	Sort          int       `form:"sort" json:"sort"`                       // comment 排序越小越往前
	Status        int       `form:"status" json:"status"`                   // comment 用户状态10是默认正常 20是禁止登陆
	Token         string    `form:"token" json:"token"`                     // comment token
	TokenTime     time.Time `form:"token_time" json:"token_time"`           // comment 登陆token时间

}
