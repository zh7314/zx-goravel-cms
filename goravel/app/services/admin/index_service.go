package admin

import (
	"errors"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"github.com/jianfengye/collection"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils"
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
	if err != nil {
		return res, err
	}
	url := facades.Storage().Url(path)

	data := make(map[string]interface{})

	data["id"] = utils.GetUniqid()
	data["src"] = url
	data["fileName"] = file.GetClientOriginalName()

	return data, nil
}

func (r *IndexService) Login(request requests.AdminLoginRequest) (res map[string]interface{}, err error) {

	if gconv.IsEmpty(request.Username) {
		return res, errors.New("账号不能为空")
	}
	if gconv.IsEmpty(request.Password) {
		return res, errors.New("密码不能为空")
	}

	var a models.Admin
	err = facades.Orm().Query().Where("name", request.Username).FirstOrFail(&a)

	if err == orm.ErrRecordNotFound {
		return res, errors.New("账号错误")
	}

	if str.Md5(str.Md5(request.Password+a.Salt)) != a.Password {
		return res, errors.New("密码错误")
	}

	token := utils.GetUniqid()

	a.TokenTime = time.Now()
	a.Token = token

	err = facades.Orm().Query().Save(&a)
	if err != nil {
		return res, err
	}

	data := make(map[string]interface{})
	userInfo := make(map[string]interface{})

	userInfo["dashboard"] = ""
	userInfo["userId"] = a.ID
	userInfo["avatar"] = a.Avatar
	userInfo["userName"] = a.RealName
	userInfo["role"] = []string{"SA", "admin", "Auditor"}

	data["token"] = token
	data["tokenTime"] = global.TOKEN_TIME
	data["tokenKey"] = global.API_TOKEN
	data["userInfo"] = userInfo

	return data, nil
}

func (r *IndexService) GetInfo(adminId int64) (res map[string]interface{}, err error) {

	if gconv.IsEmpty(adminId) {
		return res, errors.New("用户id不能为空")
	}

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

func (r *IndexService) Logout(adminId int64) (res bool, err error) {

	if gconv.IsEmpty(adminId) {
		return res, errors.New("用户id不能为空")
	}

	var a models.Admin
	err = facades.Orm().Query().Where("id", adminId).FirstOrFail(&a)
	if err != nil {
		return res, errors.New("用户不存在")
	}
	a.Token = ""

	err = facades.Orm().Query().Save(&a)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *IndexService) GetVersion() (string, error) {

	return "1.6.9", nil
}

func (r *IndexService) ChangePwd(adminId int64, request requests.AdminChangePwdRequest) (res bool, err error) {

	if gconv.IsEmpty(adminId) {
		return res, errors.New("用户id不能为空")
	}
	if gconv.IsEmpty(request.UserPassword) {
		return res, errors.New("用户id不能为空")
	}
	if gconv.IsEmpty(request.NewPassword) {
		return res, errors.New("用户id不能为空")
	}
	if gconv.IsEmpty(request.ConfirmNewPassword) {
		return res, errors.New("用户id不能为空")
	}
	if request.NewPassword != request.ConfirmNewPassword {
		return false, errors.New("两次新密码不一致")
	}

	var a models.Admin
	err = facades.Orm().Query().Where("id", adminId).FirstOrFail(&a)
	if err != nil {
		return false, errors.New("用户不存在")
	}

	if str.Md5(str.Md5(request.UserPassword+a.Salt)) != a.Password {
		return res, errors.New("密码错误")
	}

	a.Password = str.Md5(str.Md5(request.NewPassword + a.Salt))
	//踢掉token
	a.Token = ""

	err = facades.Orm().Query().Save(&a)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *IndexService) GetMenu(adminId int64) (res map[string]interface{}, err error) {

	if gconv.IsEmpty(adminId) {
		return res, errors.New("用户id不能为空")
	}

	var a models.Admin
	err = facades.Orm().Query().Where("id", adminId).FirstOrFail(&a)
	if err != nil {
		return res, errors.New("用户不存在")
	}

	data := make(map[string]interface{})
	data["dashboardGrid"] = []string{"welcome", "ver", "time", "progress", "echarts", "about"}
	data["permissions"] = []string{"list.add", "list.edit", "list.delete", "user.add", "user.edit", "user.delete"}

	menu, err := NewCommonService().GetMenu(a.ID, a.IsAdmin)
	data["menu"] = menu

	return data, nil
}

func (r *IndexService) GetGroupTree() (res []map[string]interface{}, err error) {

	var count int64
	facades.Orm().Query().Model(&models.AdminGroup{}).Count(&count)
	if count <= 0 {
		return res, errors.New("为空")
	}

	var list []*models.AdminGroup
	//索取所有的菜单
	facades.Orm().Query().Get(&list)

	result := make([]map[string]interface{}, len(list))
	for i, v := range list {
		result[i] = make(map[string]interface{})
		result[i]["id"] = v.ID
		result[i]["parent_id"] = v.ParentId
		result[i]["name"] = v.Name
		result[i]["create_at"] = v.CreateAt
		result[i]["update_at"] = v.UpdateAt
		result[i]["sort"] = v.Sort
		result[i]["permission_ids"] = v.PermissionIds
		result[i]["children"] = v.Children
	}

	menu := NewCommonService().TreeMenu(result, 0)
	if err != nil {
		return res, err
	}
	return menu, nil
}

func (r *IndexService) GetMenuTree() (res []map[string]interface{}, err error) {

	var count int64
	facades.Orm().Query().Model(&models.AdminGroup{}).Count(&count)
	if count <= 0 {
		return res, errors.New("为空")
	}

	var list []*models.AdminGroup
	//索取所有的菜单
	facades.Orm().Query().Get(&list)

	result := make([]map[string]interface{}, len(list))
	for i, v := range list {
		result[i] = make(map[string]interface{})
		result[i]["id"] = v.ID
		result[i]["parent_id"] = v.ParentId
		result[i]["name"] = v.Name
		result[i]["create_at"] = v.CreateAt
		result[i]["update_at"] = v.UpdateAt
		result[i]["sort"] = v.Sort
		result[i]["permission_ids"] = v.PermissionIds
		result[i]["children"] = v.Children
	}

	menu := NewCommonService().TreeMenu(result, 0)
	if err != nil {
		return res, err
	}
	return menu, nil
}

func (r *IndexService) GetDownloadCateTree() (res []map[string]interface{}, err error) {

	return res, nil
}

func (r *IndexService) GetNewsCateTree() (res []map[string]interface{}, err error) {

	return res, nil
}

func (r *IndexService) GetProductCateTree() (res []map[string]interface{}, err error) {

	return res, nil
}

func (r *IndexService) GetVideoCateTree() (res []map[string]interface{}, err error) {

	return res, nil
}

func (r *IndexService) GetBannerCateTree() (res []map[string]interface{}, err error) {

	return res, nil
}
