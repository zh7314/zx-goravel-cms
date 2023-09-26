package admin

import (
	"github.com/goravel/framework/contracts/http"
	requests "goravel/app/requests/admin"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type ProductCateController struct {
}

func NewProductCateController() *ProductCateController {
	return &ProductCateController{}
}

func (r *ProductCateController) GetList(ctx http.Context) http.Response {

	var request requests.ProductCateRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewProductCateService().GetList(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *ProductCateController) GetAll(ctx http.Context) http.Response {

	var request requests.ProductCateRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewProductCateService().GetAll(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *ProductCateController) GetOne(ctx http.Context) http.Response {

	var request requests.ProductCateRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewProductCateService().GetOne(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *ProductCateController) Add(ctx http.Context) http.Response {

	var request requests.ProductCateRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewProductCateService().Add(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *ProductCateController) Save(ctx http.Context) http.Response {

	var request requests.ProductCateRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewProductCateService().Save(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *ProductCateController) Delete(ctx http.Context) http.Response {

	var request requests.ProductCateRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewProductCateService().Delete(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}
