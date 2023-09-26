package admin

import (
	"github.com/goravel/framework/contracts/http"
	requests "goravel/app/requests/admin"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type SeoController struct {
}

func NewSeoController() *SeoController {
	return &SeoController{}
}

func (r *SeoController) GetList(ctx http.Context) http.Response {

	var request requests.SeoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewSeoService().GetList(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *SeoController) GetAll(ctx http.Context) http.Response {

	var request requests.SeoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewSeoService().GetAll(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *SeoController) GetOne(ctx http.Context) http.Response {

	var request requests.SeoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewSeoService().GetOne(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *SeoController) Add(ctx http.Context) http.Response {

	var request requests.SeoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewSeoService().Add(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *SeoController) Save(ctx http.Context) http.Response {

	var request requests.SeoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewSeoService().Save(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *SeoController) Delete(ctx http.Context) http.Response {

	var request requests.SeoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewSeoService().Delete(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}
