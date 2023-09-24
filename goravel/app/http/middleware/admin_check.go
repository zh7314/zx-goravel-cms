package middleware

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
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

	var admin models.Admin
	err := facades.Orm().Query().Where("token", token).FirstOrFail(&admin)
	if err != nil {
		return false, errors.New("token不存在")
	}

	//timeObj, err := utils.LocalTimeToTime(admin.TokenTime)

	//if err != nil {
	//	return false, errors.New("token时间解析错误")
	//}

	//if time.Now().Unix() > timeObj.Unix()+int64(global.TOKEN_TIME) {
	//	return false, errors.New("token过期，请重新登录")
	//}

	return true, nil
}
