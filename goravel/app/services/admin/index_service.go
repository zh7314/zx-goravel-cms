package admin

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/jianfengye/collection"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"goravel/app/utils/global"
	"goravel/app/utils/str"
	"time"
)

type IndexService struct {
}

func NewIndexService() *IndexService {
	return &IndexService{}
}

func (r *IndexService) UploadFile(file filesystem.File, acceptExt []string, fileType string) (res map[string]interface{}, err error) {

	size, err := file.Size()
	if err != nil {
		return res, err
	}
	if size <= 0 {
		return res, errors.New("上传文件大小为空")
	}

	ext := file.GetClientOriginalExtension()
	co := collection.NewStrCollection(acceptExt)

	if !co.Contains(ext) {
		return res, errors.New("上传类型不允许")
	}

	if fileType == "image" {
		//图片安全检测 todo
	}
	//本地存储
	path, err := file.Store("upload/" + fileType + "/" + time.Now().Format("20060102"))
	fmt.Println(path)
	url := facades.Storage().Url(path)
	fmt.Println(url)

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

	admin.TokenTime = time.Now()
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

func (r *IndexService) GetInfo(adminId int64) (res map[string]interface{}, err error) {

	var a models.Admin
	err = facades.Orm().Query().Where("id", adminId).FirstOrFail(&a)
	if err != nil {
		return res, errors.New("用户不存在")
	}

	data := make(map[string]interface{})
	data["name"] = a.Name
	data["avatar"] = a.Avatar
	data["introduction"] = "admin"
	data["roles"] = []string{"admin"}

	return data, nil
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
