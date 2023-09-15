package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers/admin"
	"goravel/app/http/middleware"
)

func Api() {

	//无权限
	facades.Route().Prefix("api/admin").Middleware(middleware.AdminCheck(), middleware.Recovery()).Group(func(router route.Router) {

		router.Post("login", admin.NewIndexController().Login)           //登录请求
		router.Post("uploadPic", admin.NewIndexController().UploadPic)   //上传图片文件
		router.Post("uploadFile", admin.NewIndexController().UploadFile) //上传普通文件
	})

	//有权限检查
	facades.Route().Prefix("admin").Middleware(middleware.AdminCheck(), middleware.Recovery()).Group(func(router route.Router) {

		//router.Post("getMenu", admin.NewIndexController().GetMenu)       //获取菜单信息
		//router.Post("getInfo", admin.NewIndexController().GetInfo)       //获取用户信息
		//router.Post("logout", admin.NewIndexController().logout)         //登出
		//router.Post("getVersion", admin.NewIndexController().getVersion) //获取版本信息
		//router.Post("changePwd", admin.NewIndexController().changePwd)   //修改密码
		//
		//router.Post("getGroupTree", admin.NewIndexController().GetGroupTree)
		//router.Post("getMenuTree", admin.NewIndexController().GetMenuTree)
		//router.Post("getDownloadCateTree", admin.NewIndexController().GetDownloadCateTree)
		//router.Post("getNewsCateTree", admin.NewIndexController().GetNewsCateTree)
		//router.Post("getProductCateTree", admin.NewIndexController().GetProductCateTree)
		//router.Post("getVideoCateTree", admin.NewIndexController().GetVideoCateTree)
		//router.Post("getBannerCateTree", admin.NewIndexController().GetBannerCateTree)

		router.Prefix("admin").Group(func(router1 route.Router) {
			router1.Get("/getList", admin.NewAdminController().GetList)
			router1.Get("/getAll", admin.NewAdminController().GetAll)
			router1.Get("/add", admin.NewAdminController().Add)
			router1.Get("/save", admin.NewAdminController().Save)
			router1.Get("/delete", admin.NewAdminController().Delete)
		})

		//router.Prefix("index").Group(func(router1 route.Router) {
		//	router1.Get("/getList", admin.NewIndexController().GetList)
		//	router1.Get("/getAll", admin.NewIndexController().GetAll)
		//	router1.Get("/add", admin.NewIndexController().Add)
		//	router1.Get("/save", admin.NewIndexController().Save)
		//	router1.Get("/delete", admin.NewIndexController().Delete)
		//})

	})
}
