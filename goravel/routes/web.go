package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
	"goravel/app/http/controllers/api"
	"goravel/app/http/middleware"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	//无权限
	facades.Route().Prefix("/open").Middleware(middleware.Recovery()).Group(func(router route.Router) {

		router.Get("/test", api.NewTestController().Test)
	})
}
