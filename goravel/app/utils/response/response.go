package response

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/global"
)

func Success(ctx http.Context, data interface{}, msg string) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"code": global.SUCCESS,
		"data": data,
		"msg":  msg,
	})
	return
}

func Fail(ctx http.Context, data interface{}, msg string) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"code": global.FAIL,
		"data": data,
		"msg":  msg,
	})
	return
}

func AbortFail(ctx http.Context, data interface{}, msg string) {
	ctx.Request().AbortWithStatusJson(http.StatusOK, http.Json{
		"code": global.FAIL,
		"data": data,
		"msg":  msg,
	})
	return
}

func Grant(ctx http.Context, data interface{}, msg string) {
	ctx.Response().Json(http.StatusOK, http.Json{
		"code": global.GRANT,
		"data": data,
		"msg":  msg,
	})
	return
}
