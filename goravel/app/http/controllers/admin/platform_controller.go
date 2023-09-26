package admin

import (
	"github.com/goravel/framework/contracts/http"
	requests "goravel/app/requests/admin"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type PlatformController struct {
}

func NewPlatformController() *PlatformController {
	return &PlatformController{}
}

func (r *PlatformController) GetList(ctx http.Context) http.Response {

	var request requests.PlatformRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewPlatformService().GetList(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *PlatformController) GetAll(ctx http.Context) http.Response {

	var request requests.PlatformRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewPlatformService().GetAll(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *PlatformController) GetOne(ctx http.Context) http.Response {

	var request requests.PlatformRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewPlatformService().GetOne(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *PlatformController) Add(ctx http.Context) http.Response {

	var request requests.PlatformRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewPlatformService().Add(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *PlatformController) Save(ctx http.Context) http.Response {

	var request requests.PlatformRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewPlatformService().Save(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *PlatformController) Delete(ctx http.Context) http.Response {

	var request requests.PlatformRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewPlatformService().Delete(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}
