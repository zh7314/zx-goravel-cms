package admin

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
	"strings"
)

type AdminGroupService struct {
}

func NewAdminGroupService() *AdminGroupService {
	return &AdminGroupService{}
}

func (r *AdminGroupService) GetList(request requests.AdminGroupRequest) (map[string]interface{}, error) {

	var list []*models.AdminGroup
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		orm = orm.Where("parent_id", request.ParentId)
	}
	if !gconv.IsEmpty(request.PermissionIds) {
		orm = orm.Where("permission_ids", request.PermissionIds)
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

func (r *AdminGroupService) GetAll(request requests.AdminGroupRequest) ([]*models.AdminGroup, error) {

	var list []*models.AdminGroup

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		orm = orm.Where("parent_id", request.ParentId)
	}
	if !gconv.IsEmpty(request.PermissionIds) {
		orm = orm.Where("permission_ids", request.PermissionIds)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminGroupService) GetOne(id int64) (res models.AdminGroup, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var adminGroup models.AdminGroup
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&adminGroup)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return adminGroup, nil
}

func (r *AdminGroupService) Add(request requests.AdminGroupRequest) (bool, error) {

	var adminGroup models.AdminGroup

	if !gconv.IsEmpty(request.Name) {
		adminGroup.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		adminGroup.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.PermissionIds) {
		var permissionIds = make([]string, len(request.PermissionIds))
		for k, v := range request.PermissionIds {
			permissionIds[k] = fmt.Sprintf("%d", v)
		}
		adminGroup.PermissionIds = strings.Join(permissionIds, ",")
	}
	if !gconv.IsEmpty(request.Sort) {
		adminGroup.Sort = request.Sort
	}

	err := facades.Orm().Query().Create(&adminGroup)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminGroupService) Save(request requests.AdminGroupRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var adminGroup models.AdminGroup
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&adminGroup)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Name) {
		adminGroup.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		adminGroup.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.PermissionIds) {
		var permissionIds = make([]string, len(request.PermissionIds))
		for k, v := range request.PermissionIds {
			permissionIds[k] = fmt.Sprintf("%d", v)
		}
		adminGroup.PermissionIds = strings.Join(permissionIds, ",")
	}
	if !gconv.IsEmpty(request.Sort) {
		adminGroup.Sort = request.Sort
	}

	err = facades.Orm().Query().Save(&adminGroup)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminGroupService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var adminGroup models.AdminGroup
	_, err := facades.Orm().Query().Where("id", id).Delete(&adminGroup)
	if err != nil {
		return false, err
	}
	return true, nil
}
