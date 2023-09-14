package admin

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
	"goravel/app/utils/response"
)

type ApiController struct {
	//Dependent services
}

func NewApiController() *ApiController {
	return &ApiController{
		//Inject services
	}
}

func (r *ApiController) GetToken(ctx http.Context) {

	key := ctx.Request().Input("key")
	secret := ctx.Request().Input("secret")

	data, err := services.NewApiService().GetToken(key, secret)
	if err != nil {
		response.Fail(ctx, data, err.Error())
		return
	}
	response.Success(ctx, data, "获取成功")
	return
}
