package middleware

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils"
	"goravel/app/utils/response"
)

func Recovery() http.Middleware {
	return func(ctx http.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("系统内部错误")
				response.AbortFail(ctx, "", utils.ErrorToString(r))
			}
		}()

		ctx.Request().Next()
	}
}
