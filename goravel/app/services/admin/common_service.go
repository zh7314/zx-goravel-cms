package admin

import (
	"errors"
	"fmt"
	"github.com/goravel/framework/facades"
	"github.com/jianfengye/collection"
	"goravel/app/models"
	"goravel/app/utils/gconv"
	"strings"
)

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (r *CommonService) GetAllowUrl() []string {
	return []string{"/api/admin/main", "/api/admin/login", "/api/admin/main", "/api/admin/logout", "/api/admin/upload", "/api/admin/getMenu", "/api/admin/uploadPic", "/api/admin/uploadFile", "/api/admin/getInfo"}
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

	//if IsAdmin == 10 {
	//	return menu, nil
	//} else {
	//	return r.FilterMenu(menu, adminId)
	//}

	return menu, nil
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
 * 过滤权限 占时无法实现
 */
func (r *CommonService) FilterMenu(menu []map[string]interface{}, adminId int64) (res []map[string]interface{}, err error) {

	//fmt.Print(menu)

	ids, err := r.GetAdminPermission(adminId, false)
	if err != nil {
		return res, err
	}
	fmt.Print(ids)

	newMenu := make(map[string]interface{})
	fmt.Print(newMenu)

	if gconv.IsEmpty(menu) {
		for _, v1 := range menu {
			//if gconv.IsEmpty(v1["children"]) {
			//	for _, v2 := range v1["children"] {
			//		fmt.Print(v2)
			//	}
			//}

			fmt.Print(v1)
		}
	}

	return res, nil
}

/*
 * 验证权限
 */
func (r *CommonService) Check(adminId int64, url string) (res bool, err error) {

	var admin models.Admin
	err = facades.Orm().Query().Where("id", adminId).FirstOrFail(&admin)
	if err != nil {
		return false, errors.New("用户没设置权限")
	}

	//管理员不需要处理
	if admin.IsAdmin != 10 {
		allowUrl := r.GetAllowUrl()
		//是否在通用url
		co := collection.NewStrCollection(allowUrl)
		if !co.Contains(url) {
			//不在$allow_url，再查询在授权数据里面是否有
			permissionIds, ok := r.GetAdminPermission(adminId, false)
			if ok != nil {
				return false, ok
			}

			urls := r.GetPermissionUrl(permissionIds)
			if gconv.IsEmpty(urls) {
				return false, errors.New("没有权限1")
			}

			c := collection.NewStrCollection(urls)
			if !c.Contains(url) {
				return false, errors.New("没有权限2")
			}

		}
	}
	return true, nil
}

func (r *CommonService) GetAdminPermission(adminId int64, father bool) (res []int64, err error) {

	var admin models.Admin
	err = facades.Orm().Query().Where("id", adminId).FirstOrFail(&admin)
	if err != nil {
		return res, errors.New("用户数据不存在")
	}
	if gconv.IsEmpty(admin.AdminGroupIds) {
		return res, errors.New("没有权限")
	}
	ids := strings.Split(admin.AdminGroupIds, ",")

	var groups []models.AdminGroup
	facades.Orm().Query().Select("permission_ids").Where("id in ?", ids).Get(&groups)

	if gconv.IsEmpty(groups) {
		return res, errors.New("没有权限")
	}

	var groupIds []string
	for _, v := range groups {
		temp := strings.Split(v.PermissionIds, ",")
		groupIds = append(groupIds, temp...)
	}

	co := collection.NewStrCollection(groupIds)
	cn := co.Unique()
	uniqueIds, err := cn.ToStrings()
	if err != nil {
		return res, err
	}
	//fmt.Print(uniqueIds)
	//防止数据出错
	var permissions []models.AdminPermission
	facades.Orm().Query().Where("id in ?", uniqueIds).Get(&permissions)

	var returnIds []int64
	for _, v := range permissions {
		if father {
			returnIds = append(returnIds, v.ParentId)
		} else {
			returnIds = append(returnIds, v.ID)
		}
	}
	return returnIds, nil
}

func (r *CommonService) GetPermissionUrl(ids []int64) []string {

	var permissions []models.AdminPermission
	facades.Orm().Query().Select("path").Where("id in ?", ids).Get(&permissions)

	//facades.Orm().Query().Find(&permissions, ids)
	//var urls []string
	//facades.Orm().Query().Model(&models.AdminPermission{}).Pluck("age", &urls)

	urls := []string{}
	if len(permissions) > 0 {

		for _, v := range permissions {
			urls = append(urls, v.Path)
		}
	}
	return urls
}
