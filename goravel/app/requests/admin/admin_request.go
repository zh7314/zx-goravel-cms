package admin

import "goravel/app/utils/local"

type AdminRequest struct {
	ID            int64           `form:"id" json:"id"`                           // 用户ID
	Name          string          `form:"name" json:"name"`                       // 系统管理用户名
	Password      string          `form:"password" json:"password"`               // 密码
	Salt          string          `form:"salt" json:"salt"`                       // 加密盐
	Sex           int             `form:"sex" json:"sex"`                         // 10保密20男30女
	Email         string          `form:"email" json:"email"`                     // 邮件
	Mobile        string          `form:"mobile" json:"mobile"`                   // 手机号码
	LoginIP       string          `form:"login_ip" json:"login_ip"`               // 登陆IP
	Status        int             `form:"status" json:"status"`                   // 用户状态10是默认正常 20是禁止登陆
	Avatar        string          `form:"avatar" json:"avatar"`                   // 用户头像
	RealName      string          `form:"real_name" json:"real_name"`             // 真是姓名
	TokenTime     local.LocalTime `form:"token_time" json:"token_time"`           // 登陆token时间
	AdminGroupIds string          `form:"admin_group_ids" json:"admin_group_ids"` // 用户权限组ID集合
	IsAdmin       int             `form:"is_admin" json:"is_admin"`               // 是否管理员 10是，99不是
	Sort          int             `form:"sort" json:"sort"`                       // 排序越小越往前
	CreateAt      local.LocalTime `form:"create_at" json:"create_at"`             // 创建时间
	UpdateAt      local.LocalTime `form:"update_at" json:"update_at"`             // 更新时间
	Token         string          `form:"token" json:"token"`                     // token
	Page          int             `form:"page" json:"page"`                       // 页码
	PageSize      int             `form:"pageSize" json:"pageSize"`               // 页大小
}
