package admin

import (
	requests "goravel/app/requests/admin"
)

type IndexService struct {
	//Dependent services
}

func NewIndexService() *IndexService {
	return &IndexService{
		//Inject model
	}
}

func (r *IndexService) Login(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) UploadPic(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) UploadFile(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetMenu(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetInfo(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) Logout(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetVersion(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) ChangePwd(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetGroupTree(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetMenuTree(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetDownloadCateTree(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetNewsCateTree(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetProductCateTree(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetVideoCateTree(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}

func (r *IndexService) GetBannerCateTree(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	//res := make(map[string]interface{})
	//res["list"] = list
	//res["count"] = count

	return res, nil
}
