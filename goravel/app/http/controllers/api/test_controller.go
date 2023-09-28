package api

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils"
	"goravel/app/utils/response"
	"goravel/app/utils/str"
	"strings"
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

	//fmt.Println(str.Md5(str.Md5("admin")))

	//objColl := collection.NewObjCollection([]Foo{a1, a2})
	//u1 := uuid.New()
	//u2, _ := uuid.NewRandom()
	//fmt.Println(u1.ID())
	//fmt.Println(u2.ID())

	num_arr := []int{1, 2, 3, 4}
	var str_arr = make([]string, len(num_arr))
	for k, v := range num_arr {
		str_arr[k] = fmt.Sprintf("%d", v)
	}
	var str1 = strings.Join(str_arr, ",")
	fmt.Println(str1)

	fmt.Println(utils.GetUniqid())
	return response.Success(ctx, str.Md5(str.Md5("admin")), "成功")

}
