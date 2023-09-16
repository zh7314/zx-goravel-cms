package admin

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	requests "goravel/app/requests/admin"
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

	var request requests.AdminRequest
	err := ctx.Request().Bind(&request)
	fmt.Println(err)

	data, ok := admin.NewAdminService().GetList(request)
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "获取成功")
	}
}

func (r *AdminController) GetAll(ctx http.Context) http.Response {

	var request requests.AdminRequest
	err := ctx.Request().Bind(&request)
	fmt.Println(err)

	data, ok := admin.NewAdminService().GetAll(request)
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *AdminController) Add(ctx http.Context) http.Response {

	var request requests.AdminRequest
	err := ctx.Request().Bind(&request)
	fmt.Println(err)

	data, ok := admin.NewAdminService().Add(request)
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *AdminController) Save(ctx http.Context) http.Response {

	var request requests.AdminRequest
	err := ctx.Request().Bind(&request)
	fmt.Println(err)

	data, ok := admin.NewAdminService().Save(request)
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}

func (r *AdminController) Delete(ctx http.Context) http.Response {

	var request requests.AdminRequest
	err := ctx.Request().Bind(&request)
	fmt.Println(err)

	data, ok := admin.NewAdminService().Delete(request.ID)
	if ok != nil {
		return response.Fail(ctx, "", ok.Error())
	} else {
		return response.Success(ctx, data, "失败")
	}
}
