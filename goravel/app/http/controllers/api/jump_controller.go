package api

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services"
	"goravel/app/utils/response"
)

type JumpController struct {
	//Dependent services
}

func NewJumpController() *JumpController {
	return &JumpController{
		//Inject services
	}
}

func (r *JumpController) DoJump(ctx http.Context) {

	str_id := ctx.Request().Input("id")

	jumpUrl, ok := services.NewJumpService().DoJump(str_id)
	if ok != nil {
		response.Fail(ctx, "", ok.Error())
		return
	} else {
		ctx.Response().Redirect(http.StatusMovedPermanently, jumpUrl)
		return
	}
}
