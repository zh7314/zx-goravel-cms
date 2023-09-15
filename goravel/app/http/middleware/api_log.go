package middleware

import (
	"github.com/goravel/framework/contracts/http"
)

func ApiLog() http.Middleware {
	return func(ctx http.Context) {
		//ip := utils.GetIP(ctx)
		//fmt.Println(ip)

		//ctx.WithValue("zx", "1111111111")

		ctx.Request().Next()
	}
}
