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

		router.Get("login", admin.NewIndexController().Login)           //登录请求
		router.Get("uploadPic", admin.NewIndexController().UploadPic)   //上传图片文件
		router.Get("uploadFile", admin.NewIndexController().UploadFile) //上传普通文件

		router.Get("/admin/getList", admin.NewAdminController().GetList)
		router.Get("/admin/getAll", admin.NewAdminController().GetAll)
		router.Get("/admin/add", admin.NewAdminController().Add)
		router.Get("/admin/save", admin.NewAdminController().Save)
		router.Get("/admin/delete", admin.NewAdminController().Delete)

	})

	//有权限检查
	facades.Route().Prefix("admin").Middleware(middleware.AdminCheck(), middleware.Recovery()).Group(func(router route.Router) {
		//router::post('/getMenu', [app\controller\Admin\IndexController::class, 'getMenu'])->name('获取菜单信息'); //获取菜单信息
		//Route::post('/getInfo', [app\controller\Admin\IndexController::class, 'getInfo']); //获取用户信息
		//Route::post('/logout', [app\controller\Admin\IndexController::class, 'logout']); //登出
		//Route::post('/getVersion', [app\controller\Admin\IndexController::class, 'getVersion']); //获取版本信息
		//Route::post('/changePwd', [app\controller\Admin\IndexController::class, 'changePwd']);
		//
		//Route::post('/getGroupTree', [app\controller\Admin\IndexController::class, 'getGroupTree']);//获取权限组树状结构
		//Route::post('/getMenuTree', [app\controller\Admin\IndexController::class, 'getMenuTree']);//获取所有权限树状结构
		//Route::post('/getDownloadCateTree', [app\controller\Admin\IndexController::class, 'getDownloadCateTree']);
		//Route::post('/getNewsCateTree', [app\controller\Admin\IndexController::class, 'getNewsCateTree']);
		//Route::post('/getProductCateTree', [app\controller\Admin\IndexController::class, 'getProductCateTree']);
		//Route::post('/getVideoCateTree', [app\controller\Admin\IndexController::class, 'getVideoCateTree']);
		//Route::post('/getBannerCateTree', [app\controller\Admin\IndexController::class, 'getBannerCateTree']);
	})
}
