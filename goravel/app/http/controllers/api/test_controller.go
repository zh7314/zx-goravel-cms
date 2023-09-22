package api

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/response"
	"goravel/app/utils/str"
)

type TestController struct {
	//Dependent services
}

func NewTestController() *TestController {
	return &TestController{
		//Inject services
	}
}

func (r *TestController) Test(ctx http.Context) http.Response {

	//panic("test1111111")
	//panic("test22222")
	//ctx.Response().Success().Json()

	fmt.Println(str.Md5(str.Md5("admin")))

	return response.Success(ctx, str.Md5(str.Md5("admin")), "成功")

}
