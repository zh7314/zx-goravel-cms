package api

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/services/admin"
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

	res, err := admin.NewCommonService().Check(5, "/api/admin/login1111")
	fmt.Print(res)
	fmt.Print(err)

	return response.Success(ctx, str.Md5(str.Md5("admin")), "成功")
}
