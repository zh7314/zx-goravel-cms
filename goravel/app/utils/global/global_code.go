package global

const (
	// 接口返回成功
	SUCCESS int = 200
	// 返回错误
	FAIL int = 400
	// 需要授权
	GRANT int = 401

	// 数据库 状态定义 10 正常 99删除
	NORMAL int = 10
	DELETE int = 99

	// 后台token名称
	API_TOKEN string = "X-Token"
	// token存活时间
	TOKEN_TIME int = 24 * 3600

	// 定义debug开启判断的字符
	DEBUG = "JUMP_DEBUG"
)
