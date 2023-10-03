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

	res, err := admin.NewCommonService().GetMenu(5, 99)
	fmt.Print(res)
	fmt.Print(err)

	//ids := []int64{1, 2, 3, 10}
	//tt := admin.NewCommonService().GetPermissionUrl(ids)
	//fmt.Print(tt)

	return response.Success(ctx, str.Md5(str.Md5("admin")), "成功")
}
