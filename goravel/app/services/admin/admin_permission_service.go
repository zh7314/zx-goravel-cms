package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type AdminPermissionService struct {
}

func NewAdminPermissionService() *AdminPermissionService {
	return &AdminPermissionService{}
}

func (r *AdminPermissionService) GetList(request requests.AdminPermissionRequest) (map[string]interface{}, error) {

	var list []*models.AdminPermission
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Component) {
		orm = orm.Where("component", request.Component)
	}
	if !gconv.IsEmpty(request.Hidden) {
		orm = orm.Where("hidden", request.Hidden)
	}
	if !gconv.IsEmpty(request.Icon) {
		orm = orm.Where("icon", request.Icon)
	}
	if !gconv.IsEmpty(request.IsMenu) {
		orm = orm.Where("is_menu", request.IsMenu)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		orm = orm.Where("parent_id", request.ParentId)
	}
	if !gconv.IsEmpty(request.Path) {
		orm = orm.Where("path", request.Path)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}

	if request.Page > 0 && request.PageSize > 0 {
		orm.Order("sort asc").Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)
	} else {
		orm.Order("sort asc").Order("id desc").Get(&list)
		count = int64(len(list))
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *AdminPermissionService) GetAll(request requests.AdminPermissionRequest) ([]*models.AdminPermission, error) {

	var list []*models.AdminPermission

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Component) {
		orm = orm.Where("component", request.Component)
	}
	if !gconv.IsEmpty(request.Hidden) {
		orm = orm.Where("hidden", request.Hidden)
	}
	if !gconv.IsEmpty(request.Icon) {
		orm = orm.Where("icon", request.Icon)
	}
	if !gconv.IsEmpty(request.IsMenu) {
		orm = orm.Where("is_menu", request.IsMenu)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		orm = orm.Where("parent_id", request.ParentId)
	}
	if !gconv.IsEmpty(request.Path) {
		orm = orm.Where("path", request.Path)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminPermissionService) GetOne(id int64) (res models.AdminPermission, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var adminPermission models.AdminPermission
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&adminPermission)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return adminPermission, nil
}

func (r *AdminPermissionService) Add(request requests.AdminPermissionRequest) (bool, error) {

	var adminPermission models.AdminPermission

	if !gconv.IsEmpty(request.Component) {
		adminPermission.Component = html.EscapeString(request.Component)
	}
	if !gconv.IsEmpty(request.Hidden) {
		adminPermission.Hidden = request.Hidden
	}
	if !gconv.IsEmpty(request.Icon) {
		adminPermission.Icon = html.EscapeString(request.Icon)
	}
	if !gconv.IsEmpty(request.IsMenu) {
		adminPermission.IsMenu = request.IsMenu
	}
	if !gconv.IsEmpty(request.Name) {
		adminPermission.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		adminPermission.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Path) {
		adminPermission.Path = html.EscapeString(request.Path)
	}
	if !gconv.IsEmpty(request.Sort) {
		adminPermission.Sort = request.Sort
	}

	err := facades.Orm().Query().Create(&adminPermission)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminPermissionService) Save(request requests.AdminPermissionRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var adminPermission models.AdminPermission
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&adminPermission)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Component) {
		adminPermission.Component = html.EscapeString(request.Component)
	}
	if !gconv.IsEmpty(request.Hidden) {
		adminPermission.Hidden = request.Hidden
	}
	if !gconv.IsEmpty(request.Icon) {
		adminPermission.Icon = html.EscapeString(request.Icon)
	}
	if !gconv.IsEmpty(request.IsMenu) {
		adminPermission.IsMenu = request.IsMenu
	}
	if !gconv.IsEmpty(request.Name) {
		adminPermission.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		adminPermission.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Path) {
		adminPermission.Path = html.EscapeString(request.Path)
	}
	if !gconv.IsEmpty(request.Sort) {
		adminPermission.Sort = request.Sort
	}

	err = facades.Orm().Query().Save(&adminPermission)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminPermissionService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var adminPermission models.AdminPermission
	_, err := facades.Orm().Query().Where("id", id).Delete(&adminPermission)
	if err != nil {
		return false, err
	}
	return true, nil
}
