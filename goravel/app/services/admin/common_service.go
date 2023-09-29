package admin

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (r *CommonService) GetMenu(adminId int64, IsAdmin int) (res []map[string]interface{}, err error) {

	var count int64
	facades.Orm().Query().Model(&models.AdminPermission{}).Count(&count)
	if count <= 0 {
		return res, nil
	}

	var list []*models.AdminPermission
	//索取所有的菜单
	facades.Orm().Query().Order("sort asc").Order("id desc").Get(&list)

	result := make([]map[string]interface{}, len(list))
	if len(list) > 0 {
		for i, v := range list {
			result[i] = make(map[string]interface{})
			result[i]["id"] = v.ID
			result[i]["parent_id"] = v.ParentId
			result[i]["name"] = v.Name
			result[i]["path"] = v.Path
			result[i]["component"] = v.Component
			result[i]["is_menu"] = v.IsMenu
			result[i]["icon"] = v.Icon
			result[i]["create_at"] = v.CreateAt
			result[i]["update_at"] = v.UpdateAt
			result[i]["sort"] = v.Sort
			result[i]["hidden"] = v.Hidden
			result[i]["children"] = v.Children

			meta := make(map[string]interface{})
			meta["title"] = v.Name
			meta["icon"] = v.Icon
			if v.Hidden == 10 {
				meta["hidden"] = false
			} else {
				meta["hidden"] = true
			}
			result[i]["meta"] = meta
		}
	}

	menu := r.TreeMenu(result, 0)

	if IsAdmin == 10 {
		return menu, nil
	} else {
		return r.FilterMenu(menu, adminId)
	}
}

func (r *CommonService) TreeMenu(menu []map[string]interface{}, parentId int64) []map[string]interface{} {

	data := make([]map[string]interface{}, 0)
	for _, v := range menu {
		if v["parent_id"] == parentId {
			v["children"] = r.TreeMenu(menu, v["id"].(int64))
			data = append(data, v)
		}
	}
	return data
}

// TreeMenus 不通用的递归菜单
func (r *CommonService) TreeMenus(menu []*models.AdminPermission, parentId int64) []*models.AdminPermission {

	data := make([]*models.AdminPermission, 0)
	for _, v := range menu {
		if v.ParentId == parentId {
			//v.Label = v.Name
			v.Children = r.TreeMenus(menu, v.ID)
			data = append(data, v)
		}
	}
	return data
}

/*
 * 过滤权限
 */
func (r *CommonService) FilterMenu(menu []map[string]interface{}, adminId int64) (res []map[string]interface{}, err error) {

	return res, nil
}

func (r *CommonService) GetAllowUrl() []string {
	return []string{"/api/admin/main", "/api/admin/login", "/api/admin/main", "/api/admin/logout", "/api/admin/upload", "/api/admin/adminPermission/getMenu", "/api/admin/adminPermission/getPermissionTree", "/api/admin/getInfo"}
}

/*
 * 验证权限
 */
func (r *CommonService) Check(adminId int64, url string) (res bool, err error) {

	var admin models.Admin

	_, err = facades.Orm().Query().Where("id", adminId).Delete(&admin)
	if err != nil {
		return false, errors.New("管理员信息丢失，联系管理员")
	}
	if admin.IsAdmin != 10 {
		urls := r.GetAllowUrl()
		fmt.Print(urls)

	}
	return true, nil
}
