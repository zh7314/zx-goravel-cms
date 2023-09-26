package admin

type AdminChangePwdRequest struct {
	ID                 int64  `form:"id" json:"id"`                                 // ID
	UserPassword       string `form:"userPassword" json:"userPassword"`             // comment 用户密码
	NewPassword        string `form:"newPassword" json:"newPassword"`               // comment 新密码
	ConfirmNewPassword string `form:"confirmNewPassword" json:"confirmNewPassword"` // comment 确认新密码
}
