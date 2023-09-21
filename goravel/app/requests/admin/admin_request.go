package admin

type AdminRequest struct {
	ID            int64  `form:"id" json:"id"`                           // ID
	Page          int    `form:"page" json:"page"`                       // 页码
	PageSize      int    `form:"pageSize" json:"pageSize"`               // 页大小
	Name          string `form:"name" json:"name"`                       // comment 系统管理用户名
	Password      string `form:"password" json:"password"`               // comment 密码
	Salt          string `form:"salt" json:"salt"`                       // comment 加密盐
	Sex           int    `form:"sex" json:"sex"`                         // comment 10保密20男30女
	Email         string `form:"email" json:"email"`                     // comment 邮件
	Mobile        string `form:"mobile" json:"mobile"`                   // comment 手机号码
	LoginIp       string `form:"login_ip" json:"login_ip"`               // comment 登陆IP
	Status        int    `form:"status" json:"status"`                   // comment 用户状态10是默认正常 20是禁止登陆
	Avatar        string `form:"avatar" json:"avatar"`                   // comment 用户头像
	RealName      string `form:"real_name" json:"real_name"`             // comment 真是姓名
	TokenTime     string `form:"token_time" json:"token_time"`           // comment 登陆token时间
	AdminGroupIds string `form:"admin_group_ids" json:"admin_group_ids"` // comment 用户权限组ID集合
	IsAdmin       int    `form:"is_admin" json:"is_admin"`               // comment 是否管理员 10是，99不是
	Sort          int    `form:"sort" json:"sort"`                       // comment 排序越小越往前
	Token         string `form:"token" json:"token"`                     // comment token

}
