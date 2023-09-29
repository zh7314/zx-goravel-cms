package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

func ApiLog() http.Middleware {
	return func(ctx http.Context) {

		_, err := apiLog(ctx)
		if err != nil {
			fmt.Print("系统日志写入错误", err.Error())
		}

		ctx.Request().Next()
	}
}

func apiLog(ctx http.Context) (res bool, ok error) {

	var log models.RequestLog
	log.Method = ctx.Request().Method()

	param := ctx.Request().All()
	if data, err := json.Marshal(param); err == nil {
		log.Params = string(data)
	}
	log.Url = ctx.Request().Url()
	log.Ip = ctx.Request().Ip()

	headers := ctx.Request().Headers()
	if header, err := json.Marshal(headers); err == nil {
		log.Header = string(header)
	}

	err := facades.Orm().Query().Save(&log)
	if err != nil {
		return false, err
	}

	return true, nil
}
