package middleware

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils"
	"goravel/app/utils/global"
	"goravel/app/utils/response"
)

func AdminCheck() http.Middleware {
	return func(ctx http.Context) {

		_, err := check(ctx)
		if err != nil {
			response.AbortFail(ctx, "", err.Error())
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("系统内部错误")
				response.Fail(ctx, "", utils.ErrorToString(r))
				return
			}
		}()
		ctx.Request().Next()
	}
}

func check(ctx http.Context) (res bool, ok error) {
	token := ctx.Request().Header(global.API_TOKEN, "")
	if token == "" {
		token = ctx.Request().Input(global.API_TOKEN)
	}
	if token == "" {
		return false, errors.New("token不能为空")
	}
	return true, nil
}
