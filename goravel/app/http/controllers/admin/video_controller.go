package admin

import (
	"github.com/goravel/framework/contracts/http"
	requests "goravel/app/requests/admin"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type VideoController struct {
}

func NewVideoController() *VideoController {
	return &VideoController{}
}

func (r *VideoController) GetList(ctx http.Context) http.Response {

	var request requests.VideoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewVideoService().GetList(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *VideoController) GetAll(ctx http.Context) http.Response {

	var request requests.VideoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewVideoService().GetAll(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *VideoController) GetOne(ctx http.Context) http.Response {

	var request requests.VideoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewVideoService().GetOne(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *VideoController) Add(ctx http.Context) http.Response {

	var request requests.VideoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewVideoService().Add(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *VideoController) Save(ctx http.Context) http.Response {

	var request requests.VideoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewVideoService().Save(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *VideoController) Delete(ctx http.Context) http.Response {

	var request requests.VideoRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewVideoService().Delete(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}
