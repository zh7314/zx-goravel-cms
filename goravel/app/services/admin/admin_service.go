package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type AdminService struct {
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (r *AdminService) GetList(request requests.AdminRequest) (map[string]interface{}, error) {

	var list []models.Admin
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.Salt) {
		orm = orm.Where("salt", request.Salt)
	}
	if !gconv.IsEmpty(request.Sex) {
		orm = orm.Where("sex", request.Sex)
	}
	if !gconv.IsEmpty(request.Email) {
		orm = orm.Where("email", request.Email)
	}
	if !gconv.IsEmpty(request.Mobile) {
		orm = orm.Where("mobile", request.Mobile)
	}
	if !gconv.IsEmpty(request.LoginIp) {
		orm = orm.Where("login_ip", request.LoginIp)
	}
	if !gconv.IsEmpty(request.Status) {
		orm = orm.Where("status", request.Status)
	}
	if !gconv.IsEmpty(request.Avatar) {
		orm = orm.Where("avatar", request.Avatar)
	}
	if !gconv.IsEmpty(request.RealName) {
		orm = orm.Where("real_name", request.RealName)
	}
	if !gconv.IsEmpty(request.AdminGroupIds) {
		orm = orm.Where("admin_group_ids", request.AdminGroupIds)
	}
	if !gconv.IsEmpty(request.IsAdmin) {
		orm = orm.Where("is_admin", request.IsAdmin)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}

	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *AdminService) GetAll(request requests.AdminRequest) ([]models.Admin, error) {

	var list []models.Admin

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.Password) {
		orm = orm.Where("password", request.Password)
	}
	if !gconv.IsEmpty(request.Salt) {
		orm = orm.Where("salt", request.Salt)
	}
	if !gconv.IsEmpty(request.Sex) {
		orm = orm.Where("sex", request.Sex)
	}
	if !gconv.IsEmpty(request.Email) {
		orm = orm.Where("email", request.Email)
	}
	if !gconv.IsEmpty(request.Mobile) {
		orm = orm.Where("mobile", request.Mobile)
	}
	if !gconv.IsEmpty(request.LoginIp) {
		orm = orm.Where("login_ip", request.LoginIp)
	}
	if !gconv.IsEmpty(request.Status) {
		orm = orm.Where("status", request.Status)
	}
	if !gconv.IsEmpty(request.Avatar) {
		orm = orm.Where("avatar", request.Avatar)
	}
	if !gconv.IsEmpty(request.RealName) {
		orm = orm.Where("real_name", request.RealName)
	}
	if !gconv.IsEmpty(request.TokenTime) {
		orm = orm.Where("token_time", request.TokenTime)
	}
	if !gconv.IsEmpty(request.AdminGroupIds) {
		orm = orm.Where("admin_group_ids", request.AdminGroupIds)
	}
	if !gconv.IsEmpty(request.IsAdmin) {
		orm = orm.Where("is_admin", request.IsAdmin)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Token) {
		orm = orm.Where("token", request.Token)
	}

	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminService) GetOne(id int64) (res models.Admin, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var admin models.Admin
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&admin)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return admin, nil
}

func (r *AdminService) Add(request requests.AdminRequest) (bool, error) {

	var admin models.Admin

	admin.Name = html.EscapeString(request.Name)
	admin.Password = html.EscapeString(request.Password)
	admin.Salt = html.EscapeString(request.Salt)
	admin.Sex = request.Sex
	admin.Email = html.EscapeString(request.Email)
	admin.Mobile = html.EscapeString(request.Mobile)
	admin.LoginIp = html.EscapeString(request.LoginIp)
	admin.Status = request.Status
	admin.Avatar = html.EscapeString(request.Avatar)
	admin.RealName = html.EscapeString(request.RealName)
	admin.TokenTime = request.TokenTime
	admin.AdminGroupIds = html.EscapeString(request.AdminGroupIds)
	admin.IsAdmin = request.IsAdmin
	admin.Sort = request.Sort
	admin.Token = html.EscapeString(request.Token)

	err := facades.Orm().Query().Create(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminService) Save(request requests.AdminRequest) (bool, error) {

	var admin models.Admin

	admin.ID = request.ID
	admin.Name = html.EscapeString(request.Name)
	admin.Password = html.EscapeString(request.Password)
	admin.Salt = html.EscapeString(request.Salt)
	admin.Sex = request.Sex
	admin.Email = html.EscapeString(request.Email)
	admin.Mobile = html.EscapeString(request.Mobile)
	admin.LoginIp = html.EscapeString(request.LoginIp)
	admin.Status = request.Status
	admin.Avatar = html.EscapeString(request.Avatar)
	admin.RealName = html.EscapeString(request.RealName)
	admin.TokenTime = request.TokenTime
	admin.AdminGroupIds = html.EscapeString(request.AdminGroupIds)
	admin.IsAdmin = request.IsAdmin
	admin.Sort = request.Sort
	admin.Token = html.EscapeString(request.Token)

	err := facades.Orm().Query().Save(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminService) Delete(id int64) (bool, error) {

	var admin models.Admin
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
