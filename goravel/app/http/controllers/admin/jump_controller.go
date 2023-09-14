package admin

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

func (r *JumpController) GetData(ctx http.Context) {

	//fmt.Println(ctx.Value("zx"))
	jumpService := services.NewJumpService()
	data := jumpService.GetData()

	response.Success(ctx, data, "获取成功")
	return
}

func (r *JumpController) AddLink(ctx http.Context) {

	url := ctx.Request().Input("url")
	end_time := ctx.Request().Input("end_time")

	//jumpService := services.NewJumpService()
	data, ok := services.NewJumpService().AddLink(url, end_time)
	if ok != nil {
		response.Fail(ctx, "", ok.Error())
		return
	} else {
		response.Success(ctx, data, "增加成功")
		return
	}
}
