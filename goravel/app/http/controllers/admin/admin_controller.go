package admin

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type AdminController struct {
	//Dependent services
}

func NewAdminController() *AdminController {
	return &AdminController{
		//Inject services
	}
}

func (r *AdminController) GetList(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().GetList()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *AdminController) GetAll(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().GetAll()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *AdminController) Add(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().Add()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *AdminController) Save(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().Save()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *AdminController) Delete(ctx http.Context) http.Response {

	data, ok := admin.NewAdminService().Delete()
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}
