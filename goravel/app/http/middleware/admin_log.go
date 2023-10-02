package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	"goravel/app/utils/gconv"
)

func AdminLog() http.Middleware {
	return func(ctx http.Context) {

		_, err := log(ctx)
		if err != nil {
			fmt.Print("系统日志写入错误", err.Error())
		}

		ctx.Request().Next()
	}
}

func log(ctx http.Context) (res bool, ok error) {

	param := ctx.Request().All()

	var log models.AdminLog
	log.Method = ctx.Request().Method()

	if data, err := json.Marshal(param); err == nil {
		log.Data = string(data)
	}
	log.Url = ctx.Request().Url()
	log.Path = ctx.Request().Path()

	adminId := ctx.Value("admin_id")
	if !gconv.IsEmpty(adminId) {
		log.AdminId = adminId.(int64)
	}

	log.RequestIp = ctx.Request().Ip()

	err := facades.Orm().Query().Save(&log)
	if err != nil {
		return false, err
	}

	return true, nil
}
