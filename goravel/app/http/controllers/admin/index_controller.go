package admin

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type IndexController struct {
	//Dependent services
}

func NewIndexController() *IndexController {
	return &IndexController{
		//Inject services
	}
}

func (r *IndexController) Login(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *IndexController) UploadPic(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *IndexController) UploadFile(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func (r *IndexController) GetList(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().GetList()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *IndexController) GetAll(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().GetAll()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *IndexController) Add(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().Add()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *IndexController) Save(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().Save()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *IndexController) Delete(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().Delete()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}
