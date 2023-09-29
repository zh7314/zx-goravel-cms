package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers/admin"
	"goravel/app/http/middleware"
)

func Api() {

	//无权限
	facades.Route().Prefix("api/admin").Middleware(middleware.AdminLog(), middleware.Recovery()).Group(func(router route.Router) {

		router.Post("login", admin.NewIndexController().Login)           //登录请求
		router.Post("uploadPic", admin.NewIndexController().UploadPic)   //上传图片文件
		router.Post("uploadFile", admin.NewIndexController().UploadFile) //上传普通文件
	})

	//有权限检查 middleware.AdminCheck()
	facades.Route().Prefix("api/admin").Middleware(middleware.AdminCheck(), middleware.AdminLog(), middleware.Recovery()).Group(func(router route.Router) {

		router.Post("getMenu", admin.NewIndexController().GetMenu)       //获取菜单信息
		router.Post("getInfo", admin.NewIndexController().GetInfo)       //获取用户信息
		router.Post("logout", admin.NewIndexController().Logout)         //登出
		router.Post("getVersion", admin.NewIndexController().GetVersion) //获取版本信息
		router.Post("changePwd", admin.NewIndexController().ChangePwd)   //修改密码

		router.Post("getGroupTree", admin.NewIndexController().GetGroupTree)
		router.Post("getMenuTree", admin.NewIndexController().GetMenuTree)
		router.Post("getDownloadCateTree", admin.NewIndexController().GetDownloadCateTree)
		router.Post("getNewsCateTree", admin.NewIndexController().GetNewsCateTree)
		router.Post("getProductCateTree", admin.NewIndexController().GetProductCateTree)
		router.Post("getVideoCateTree", admin.NewIndexController().GetVideoCateTree)
		router.Post("getBannerCateTree", admin.NewIndexController().GetBannerCateTree)

		router.Prefix("admin").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewAdminController().GetList)
			router1.Post("/getAll", admin.NewAdminController().GetAll)
			router1.Post("/getOne", admin.NewAdminController().GetOne)
			router1.Post("/add", admin.NewAdminController().Add)
			router1.Post("/save", admin.NewAdminController().Save)
			router1.Post("/delete", admin.NewAdminController().Delete)
		})

		router.Prefix("adminGroup").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewAdminGroupController().GetList)
			router1.Post("/getAll", admin.NewAdminGroupController().GetAll)
			router1.Post("/getOne", admin.NewAdminGroupController().GetOne)
			router1.Post("/add", admin.NewAdminGroupController().Add)
			router1.Post("/save", admin.NewAdminGroupController().Save)
			router1.Post("/delete", admin.NewAdminGroupController().Delete)
		})

		router.Prefix("adminLog").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewAdminLogController().GetList)
			router1.Post("/getAll", admin.NewAdminLogController().GetAll)
			router1.Post("/getOne", admin.NewAdminLogController().GetOne)
			router1.Post("/add", admin.NewAdminLogController().Add)
			router1.Post("/save", admin.NewAdminLogController().Save)
			router1.Post("/delete", admin.NewAdminLogController().Delete)
		})

		router.Prefix("adminPermission").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewAdminPermissionController().GetList)
			router1.Post("/getAll", admin.NewAdminPermissionController().GetAll)
			router1.Post("/getOne", admin.NewAdminPermissionController().GetOne)
			router1.Post("/add", admin.NewAdminPermissionController().Add)
			router1.Post("/save", admin.NewAdminPermissionController().Save)
			router1.Post("/delete", admin.NewAdminPermissionController().Delete)
		})

		router.Prefix("banner").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewBannerController().GetList)
			router1.Post("/getAll", admin.NewBannerController().GetAll)
			router1.Post("/getOne", admin.NewBannerController().GetOne)
			router1.Post("/add", admin.NewBannerController().Add)
			router1.Post("/save", admin.NewBannerController().Save)
			router1.Post("/delete", admin.NewBannerController().Delete)
		})

		router.Prefix("bannerCate").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewBannerCateController().GetList)
			router1.Post("/getAll", admin.NewBannerCateController().GetAll)
			router1.Post("/getOne", admin.NewBannerCateController().GetOne)
			router1.Post("/add", admin.NewBannerCateController().Add)
			router1.Post("/save", admin.NewBannerCateController().Save)
			router1.Post("/delete", admin.NewBannerCateController().Delete)
		})

		router.Prefix("config").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewConfigController().GetList)
			router1.Post("/getAll", admin.NewConfigController().GetAll)
			router1.Post("/getOne", admin.NewConfigController().GetOne)
			router1.Post("/add", admin.NewConfigController().Add)
			router1.Post("/save", admin.NewConfigController().Save)
			router1.Post("/delete", admin.NewConfigController().Delete)
		})

		router.Prefix("download").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewDownloadController().GetList)
			router1.Post("/getAll", admin.NewDownloadController().GetAll)
			router1.Post("/getOne", admin.NewDownloadController().GetOne)
			router1.Post("/add", admin.NewDownloadController().Add)
			router1.Post("/save", admin.NewDownloadController().Save)
			router1.Post("/delete", admin.NewDownloadController().Delete)
		})

		router.Prefix("downloadCate").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewDownloadCateController().GetList)
			router1.Post("/getAll", admin.NewDownloadCateController().GetAll)
			router1.Post("/getOne", admin.NewDownloadCateController().GetOne)
			router1.Post("/add", admin.NewDownloadCateController().Add)
			router1.Post("/save", admin.NewDownloadCateController().Save)
			router1.Post("/delete", admin.NewDownloadCateController().Delete)
		})

		router.Prefix("feedback").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewFeedbackController().GetList)
			router1.Post("/getAll", admin.NewFeedbackController().GetAll)
			router1.Post("/getOne", admin.NewFeedbackController().GetOne)
			router1.Post("/add", admin.NewFeedbackController().Add)
			router1.Post("/save", admin.NewFeedbackController().Save)
			router1.Post("/delete", admin.NewFeedbackController().Delete)
		})

		router.Prefix("file").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewFileController().GetList)
			router1.Post("/getAll", admin.NewFileController().GetAll)
			router1.Post("/getOne", admin.NewFileController().GetOne)
			router1.Post("/add", admin.NewFileController().Add)
			router1.Post("/save", admin.NewFileController().Save)
			router1.Post("/delete", admin.NewFileController().Delete)
		})

		router.Prefix("friendLink").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewFriendLinkController().GetList)
			router1.Post("/getAll", admin.NewFriendLinkController().GetAll)
			router1.Post("/getOne", admin.NewFriendLinkController().GetOne)
			router1.Post("/add", admin.NewFriendLinkController().Add)
			router1.Post("/save", admin.NewFriendLinkController().Save)
			router1.Post("/delete", admin.NewFriendLinkController().Delete)
		})

		router.Prefix("jobOffers").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewJobOffersController().GetList)
			router1.Post("/getAll", admin.NewJobOffersController().GetAll)
			router1.Post("/getOne", admin.NewJobOffersController().GetOne)
			router1.Post("/add", admin.NewJobOffersController().Add)
			router1.Post("/save", admin.NewJobOffersController().Save)
			router1.Post("/delete", admin.NewJobOffersController().Delete)
		})

		router.Prefix("lang").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewLangController().GetList)
			router1.Post("/getAll", admin.NewLangController().GetAll)
			router1.Post("/getOne", admin.NewLangController().GetOne)
			router1.Post("/add", admin.NewLangController().Add)
			router1.Post("/save", admin.NewLangController().Save)
			router1.Post("/delete", admin.NewLangController().Delete)
		})

		router.Prefix("message").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewMessageController().GetList)
			router1.Post("/getAll", admin.NewMessageController().GetAll)
			router1.Post("/getOne", admin.NewMessageController().GetOne)
			router1.Post("/add", admin.NewMessageController().Add)
			router1.Post("/save", admin.NewMessageController().Save)
			router1.Post("/delete", admin.NewMessageController().Delete)
		})

		router.Prefix("news").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewNewsController().GetList)
			router1.Post("/getAll", admin.NewNewsController().GetAll)
			router1.Post("/getOne", admin.NewNewsController().GetOne)
			router1.Post("/add", admin.NewNewsController().Add)
			router1.Post("/save", admin.NewNewsController().Save)
			router1.Post("/delete", admin.NewNewsController().Delete)
		})

		router.Prefix("newsCate").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewNewsCateController().GetList)
			router1.Post("/getAll", admin.NewNewsCateController().GetAll)
			router1.Post("/getOne", admin.NewNewsCateController().GetOne)
			router1.Post("/add", admin.NewNewsCateController().Add)
			router1.Post("/save", admin.NewNewsCateController().Save)
			router1.Post("/delete", admin.NewNewsCateController().Delete)
		})

		router.Prefix("platform").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewPlatformController().GetList)
			router1.Post("/getAll", admin.NewPlatformController().GetAll)
			router1.Post("/getOne", admin.NewPlatformController().GetOne)
			router1.Post("/add", admin.NewPlatformController().Add)
			router1.Post("/save", admin.NewPlatformController().Save)
			router1.Post("/delete", admin.NewPlatformController().Delete)
		})

		router.Prefix("product").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewProductController().GetList)
			router1.Post("/getAll", admin.NewProductController().GetAll)
			router1.Post("/getOne", admin.NewProductController().GetOne)
			router1.Post("/add", admin.NewProductController().Add)
			router1.Post("/save", admin.NewProductController().Save)
			router1.Post("/delete", admin.NewProductController().Delete)
		})

		router.Prefix("productCate").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewProductCateController().GetList)
			router1.Post("/getAll", admin.NewProductCateController().GetAll)
			router1.Post("/getOne", admin.NewProductCateController().GetOne)
			router1.Post("/add", admin.NewProductCateController().Add)
			router1.Post("/save", admin.NewProductCateController().Save)
			router1.Post("/delete", admin.NewProductCateController().Delete)
		})

		router.Prefix("requestLog").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewRequestLogController().GetList)
			router1.Post("/getAll", admin.NewRequestLogController().GetAll)
			router1.Post("/getOne", admin.NewRequestLogController().GetOne)
			router1.Post("/add", admin.NewRequestLogController().Add)
			router1.Post("/save", admin.NewRequestLogController().Save)
			router1.Post("/delete", admin.NewRequestLogController().Delete)
		})

		router.Prefix("seo").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewSeoController().GetList)
			router1.Post("/getAll", admin.NewSeoController().GetAll)
			router1.Post("/getOne", admin.NewSeoController().GetOne)
			router1.Post("/add", admin.NewSeoController().Add)
			router1.Post("/save", admin.NewSeoController().Save)
			router1.Post("/delete", admin.NewSeoController().Delete)
		})

		router.Prefix("video").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewVideoController().GetList)
			router1.Post("/getAll", admin.NewVideoController().GetAll)
			router1.Post("/getOne", admin.NewVideoController().GetOne)
			router1.Post("/add", admin.NewVideoController().Add)
			router1.Post("/save", admin.NewVideoController().Save)
			router1.Post("/delete", admin.NewVideoController().Delete)
		})

		router.Prefix("videoCate").Group(func(router1 route.Router) {
			router1.Post("/getList", admin.NewVideoCateController().GetList)
			router1.Post("/getAll", admin.NewVideoCateController().GetAll)
			router1.Post("/getOne", admin.NewVideoCateController().GetOne)
			router1.Post("/add", admin.NewVideoCateController().Add)
			router1.Post("/save", admin.NewVideoCateController().Save)
			router1.Post("/delete", admin.NewVideoCateController().Delete)
		})

	})
}
