package admin

import (
	"errors"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils"
	"goravel/app/utils/gconv"
	"goravel/app/utils/global"
	"goravel/app/utils/str"
	"time"
)

type IndexService struct {
	//Dependent services
}

func NewIndexService() *IndexService {
	return &IndexService{
		//Inject model
	}
}

func (r *IndexService) UploadPic(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	return res, nil
}

func (r *IndexService) UploadFile(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	return res, nil
}

func (r *IndexService) Login(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	if gconv.IsEmpty(request.Username) {
		return res, errors.New("账号不能为空")
	}
	if gconv.IsEmpty(request.Password) {
		return res, errors.New("密码不能为空")
	}

	var admin models.Admin
	err = facades.Orm().Query().Where("name", request.Username).FirstOrFail(&admin)

	if err == orm.ErrRecordNotFound {
		return res, errors.New("账号错误")
	}

	if str.Md5(str.Md5(request.Password+admin.Salt)) != admin.Password {
		return res, errors.New("密码错误")
	}

	token := str.GetRandomString(32)

	tokenTime, err := utils.StrToLocalTime(time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return res, errors.New("获取时间失败")
	}

	admin.TokenTime = tokenTime
	admin.Token = token

	err = facades.Orm().Query().Save(&admin)
	if err != nil {
		return res, err
	}

	data := make(map[string]interface{})
	userInfo := make(map[string]interface{})

	userInfo["dashboard"] = ""
	userInfo["userId"] = admin.ID
	userInfo["avatar"] = admin.Avatar
	userInfo["userName"] = admin.RealName
	userInfo["role"] = []string{"SA", "admin", "Auditor"}

	data["token"] = token
	data["tokenTime"] = global.TOKEN_TIME
	data["tokenKey"] = global.API_TOKEN
	data["userInfo"] = userInfo

	return data, nil
}

func (r *IndexService) GetInfo(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

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
