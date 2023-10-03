package middleware

import (
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	"goravel/app/services/admin"
	"goravel/app/utils/global"
	"goravel/app/utils/response"
	"time"
)

func AdminCheck() http.Middleware {
	return func(ctx http.Context) {

		_, err := check(ctx)
		if err != nil {
			response.Grant(ctx, "", err.Error())
		}
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

	var a models.Admin
	err := facades.Orm().Query().Where("token", token).FirstOrFail(&a)
	if err != nil {
		return false, errors.New("token不存在")
	}

	if time.Now().Unix() > a.TokenTime.Unix()+int64(global.TOKEN_TIME) {
		return false, errors.New("token过期，请重新登录")
	}

	ctx.WithValue("admin_id", a.ID)

	//权限验证
	url := ctx.Request().Path()
	_, err = admin.NewCommonService().Check(a.ID, url)
	if err != nil {
		return false, err
	}

	return true, nil
}
