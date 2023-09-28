package api

import (
	"fmt"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/response"
	"goravel/app/utils/str"
	"html"
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

	s := "<p>2132131</p>"

	s1 := html.EscapeString(s)
	fmt.Println(s1)

	s2 := html.UnescapeString(s1)
	fmt.Println(s2)
	
	return response.Success(ctx, str.Md5(str.Md5("admin")), "成功")

}
