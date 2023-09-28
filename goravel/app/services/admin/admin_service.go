package admin

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"goravel/app/utils/str"
	"html"
	"strings"
)

type AdminService struct {
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (r *AdminService) GetList(request requests.AdminRequest) (map[string]interface{}, error) {

	var list []*models.Admin
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminGroupIds) {
		orm = orm.Where("admin_group_ids", request.AdminGroupIds)
	}
	if !gconv.IsEmpty(request.Avatar) {
		orm = orm.Where("avatar", request.Avatar)
	}
	if !gconv.IsEmpty(request.Email) {
		orm = orm.Where("email", request.Email)
	}
	if !gconv.IsEmpty(request.IsAdmin) {
		orm = orm.Where("is_admin", request.IsAdmin)
	}
	if !gconv.IsEmpty(request.LoginIp) {
		orm = orm.Where("login_ip", request.LoginIp)
	}
	if !gconv.IsEmpty(request.Mobile) {
		orm = orm.Where("mobile", request.Mobile)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.Password) {
		orm = orm.Where("password", request.Password)
	}
	if !gconv.IsEmpty(request.RealName) {
		orm = orm.Where("real_name", request.RealName)
	}
	if !gconv.IsEmpty(request.Salt) {
		orm = orm.Where("salt", request.Salt)
	}
	if !gconv.IsEmpty(request.Sex) {
		orm = orm.Where("sex", request.Sex)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Status) {
		orm = orm.Where("status", request.Status)
	}
	if !gconv.IsEmpty(request.Token) {
		orm = orm.Where("token", request.Token)
	}
	if !request.TokenTime.IsZero() {
		orm = orm.Where("token_time", request.TokenTime)
	}

	if request.Page > 0 && request.PageSize > 0 {
		orm.Order("sort asc").Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)
	} else {
		orm.Order("sort asc").Order("id desc").Get(&list)
		count = int64(len(list))
	}

	result := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			result[i] = make(map[string]interface{})
			result[i]["id"] = v.ID
			result[i]["admin_group_ids"] = strings.Split(v.AdminGroupIds, ",")
			result[i]["avatar"] = v.Avatar
			result[i]["email"] = v.Email
			result[i]["is_admin"] = v.IsAdmin
			result[i]["login_ip"] = v.LoginIp
			result[i]["mobile"] = v.Mobile
			result[i]["name"] = v.Name
			result[i]["password"] = v.Password
			result[i]["real_name"] = v.RealName
			result[i]["salt"] = v.Salt
			result[i]["sex"] = v.Sex
			result[i]["sort"] = v.Sort
			result[i]["status"] = v.Status
			result[i]["token"] = v.Token
			result[i]["token_time"] = v.TokenTime
			result[i]["create_at"] = v.CreateAt
			result[i]["update_at"] = v.UpdateAt
		}
	}

	res := make(map[string]interface{})
	res["list"] = result
	res["count"] = count

	return res, nil
}

func (r *AdminService) GetAll(request requests.AdminRequest) ([]*models.Admin, error) {

	var list []*models.Admin

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminGroupIds) {
		orm = orm.Where("admin_group_ids", request.AdminGroupIds)
	}
	if !gconv.IsEmpty(request.Avatar) {
		orm = orm.Where("avatar", request.Avatar)
	}
	if !gconv.IsEmpty(request.Email) {
		orm = orm.Where("email", request.Email)
	}
	if !gconv.IsEmpty(request.IsAdmin) {
		orm = orm.Where("is_admin", request.IsAdmin)
	}
	if !gconv.IsEmpty(request.LoginIp) {
		orm = orm.Where("login_ip", request.LoginIp)
	}
	if !gconv.IsEmpty(request.Mobile) {
		orm = orm.Where("mobile", request.Mobile)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.Password) {
		orm = orm.Where("password", request.Password)
	}
	if !gconv.IsEmpty(request.RealName) {
		orm = orm.Where("real_name", request.RealName)
	}
	if !gconv.IsEmpty(request.Salt) {
		orm = orm.Where("salt", request.Salt)
	}
	if !gconv.IsEmpty(request.Sex) {
		orm = orm.Where("sex", request.Sex)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Status) {
		orm = orm.Where("status", request.Status)
	}
	if !gconv.IsEmpty(request.Token) {
		orm = orm.Where("token", request.Token)
	}
	if !request.TokenTime.IsZero() {
		orm = orm.Where("token_time", request.TokenTime)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

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

	if !gconv.IsEmpty(request.AdminGroupIds) {
		var adminGroupIds = make([]string, len(request.AdminGroupIds))
		for k, v := range request.AdminGroupIds {
			adminGroupIds[k] = fmt.Sprintf("%d", v)
		}

		admin.AdminGroupIds = strings.Join(adminGroupIds, ",")
	}
	if !gconv.IsEmpty(request.Avatar) {
		admin.Avatar = html.EscapeString(request.Avatar)
	}
	if !gconv.IsEmpty(request.Email) {
		admin.Email = html.EscapeString(request.Email)
	}
	if !gconv.IsEmpty(request.IsAdmin) {
		admin.IsAdmin = request.IsAdmin
	}
	if !gconv.IsEmpty(request.LoginIp) {
		admin.LoginIp = html.EscapeString(request.LoginIp)
	}
	if !gconv.IsEmpty(request.Mobile) {
		admin.Mobile = html.EscapeString(request.Mobile)
	}
	if !gconv.IsEmpty(request.Name) {
		admin.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Password) {
		admin.Password = str.Md5(str.Md5(html.EscapeString(request.Password)))
	}
	if !gconv.IsEmpty(request.RealName) {
		admin.RealName = html.EscapeString(request.RealName)
	}
	if !gconv.IsEmpty(request.Salt) {
		admin.Salt = html.EscapeString(request.Salt)
	}
	if !gconv.IsEmpty(request.Sex) {
		admin.Sex = request.Sex
	}
	if !gconv.IsEmpty(request.Sort) {
		admin.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Status) {
		admin.Status = request.Status
	}
	if !gconv.IsEmpty(request.Token) {
		admin.Token = html.EscapeString(request.Token)
	}
	if !request.TokenTime.IsZero() {
		admin.TokenTime = &request.TokenTime
	}

	err := facades.Orm().Query().Create(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminService) Save(request requests.AdminRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var admin models.Admin
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&admin)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.AdminGroupIds) {
		var adminGroupIds = make([]string, len(request.AdminGroupIds))
		for k, v := range request.AdminGroupIds {
			adminGroupIds[k] = fmt.Sprintf("%d", v)
		}

		admin.AdminGroupIds = strings.Join(adminGroupIds, ",")
	}
	if !gconv.IsEmpty(request.Avatar) {
		admin.Avatar = html.EscapeString(request.Avatar)
	}
	if !gconv.IsEmpty(request.Email) {
		admin.Email = html.EscapeString(request.Email)
	}
	if !gconv.IsEmpty(request.IsAdmin) {
		admin.IsAdmin = request.IsAdmin
	}
	if !gconv.IsEmpty(request.LoginIp) {
		admin.LoginIp = html.EscapeString(request.LoginIp)
	}
	if !gconv.IsEmpty(request.Mobile) {
		admin.Mobile = html.EscapeString(request.Mobile)
	}
	if !gconv.IsEmpty(request.Name) {
		admin.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Password) {
		admin.Password = str.Md5(str.Md5(html.EscapeString(request.Password)))
	}
	if !gconv.IsEmpty(request.RealName) {
		admin.RealName = html.EscapeString(request.RealName)
	}
	if !gconv.IsEmpty(request.Salt) {
		admin.Salt = html.EscapeString(request.Salt)
	}
	if !gconv.IsEmpty(request.Sex) {
		admin.Sex = request.Sex
	}
	if !gconv.IsEmpty(request.Sort) {
		admin.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Status) {
		admin.Status = request.Status
	}
	if !gconv.IsEmpty(request.Token) {
		admin.Token = html.EscapeString(request.Token)
	}
	if !request.TokenTime.IsZero() {
		admin.TokenTime = &request.TokenTime
	}

	err = facades.Orm().Query().Save(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var admin models.Admin
	_, err := facades.Orm().Query().Where("id", id).Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
