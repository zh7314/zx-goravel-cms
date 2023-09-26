package admin

import (
	"github.com/goravel/framework/contracts/http"
	requests "goravel/app/requests/admin"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type FeedbackController struct {
}

func NewFeedbackController() *FeedbackController {
	return &FeedbackController{}
}

func (r *FeedbackController) GetList(ctx http.Context) http.Response {

	var request requests.FeedbackRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewFeedbackService().GetList(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *FeedbackController) GetAll(ctx http.Context) http.Response {

	var request requests.FeedbackRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewFeedbackService().GetAll(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *FeedbackController) GetOne(ctx http.Context) http.Response {

	var request requests.FeedbackRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewFeedbackService().GetOne(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *FeedbackController) Add(ctx http.Context) http.Response {

	var request requests.FeedbackRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewFeedbackService().Add(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *FeedbackController) Save(ctx http.Context) http.Response {

	var request requests.FeedbackRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewFeedbackService().Save(request)
	if ok == nil {
    	return response.Success(ctx, data, "成功")
    } else {
    	return response.Fail(ctx, "", ok.Error())
    }
}

func (r *FeedbackController) Delete(ctx http.Context) http.Response {

	var request requests.FeedbackRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewFeedbackService().Delete(request.ID)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}
