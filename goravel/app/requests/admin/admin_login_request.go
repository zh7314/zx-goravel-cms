package admin

type AdminLoginRequest struct {
	Username   string `form:"username" json:"username"`       // comment 登录账号
	Password   string `form:"password" json:"password"`       // comment 请求url带参数
	Code       string `form:"code" json:"code"`               // comment 验证码
	CaptchaKey string `form:"captcha_key" json:"captcha_key"` // comment 验证码标识
}
