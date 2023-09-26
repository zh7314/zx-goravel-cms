package admin

import (
	"github.com/goravel/framework/contracts/http"
	requests "goravel/app/requests/admin"
	"goravel/app/services/admin"
	"goravel/app/utils/response"
)

type IndexController struct {
	//Dependent services
}

func NewIndexController() *IndexController {
	return &IndexController{
		//Inject services
	}
}

func (r *IndexController) Login(ctx http.Context) http.Response {

	var request requests.AdminLoginRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewIndexService().Login(request)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) UploadPic(ctx http.Context) http.Response {

	file, err := ctx.Request().File("file")
	if err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewIndexService().UploadFile(file, []string{"jpg", "jpeg", "png", "mbp", "gif"}, "image")
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) UploadFile(ctx http.Context) http.Response {

	file, err := ctx.Request().File("file")
	if err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	data, ok := admin.NewIndexService().UploadFile(file, []string{"xls", "xlsx", "pdf", "xls", "xlsx", "doc", "docx", "ppt", "zip", "pptx", "mp4", "flv"}, "file")
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetInfo(ctx http.Context) http.Response {

	adminId := ctx.Value("admin_id").(int64)

	data, ok := admin.NewIndexService().GetInfo(adminId)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) Logout(ctx http.Context) http.Response {

	adminId := ctx.Value("admin_id").(int64)

	data, ok := admin.NewIndexService().Logout(adminId)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetVersion(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetVersion()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) ChangePwd(ctx http.Context) http.Response {

	var request requests.AdminChangePwdRequest
	if err := ctx.Request().Bind(&request); err != nil {
		return response.Fail(ctx, "", err.Error())
	}

	adminId := ctx.Value("admin_id").(int64)

	data, ok := admin.NewIndexService().ChangePwd(adminId, request)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetMenu(ctx http.Context) http.Response {

	adminId := ctx.Value("admin_id").(int64)

	data, ok := admin.NewIndexService().GetMenu(adminId)
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetGroupTree(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetGroupTree()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetMenuTree(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetMenuTree()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetDownloadCateTree(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetDownloadCateTree()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetNewsCateTree(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetNewsCateTree()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetProductCateTree(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetProductCateTree()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetVideoCateTree(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetVideoCateTree()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}

func (r *IndexController) GetBannerCateTree(ctx http.Context) http.Response {

	data, ok := admin.NewIndexService().GetBannerCateTree()
	if ok == nil {
		return response.Success(ctx, data, "成功")
	} else {
		return response.Fail(ctx, "", ok.Error())
	}
}
