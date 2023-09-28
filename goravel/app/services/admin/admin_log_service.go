package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type AdminLogService struct {
}

func NewAdminLogService() *AdminLogService {
	return &AdminLogService{}
}

func (r *AdminLogService) GetList(request requests.AdminLogRequest) (map[string]interface{}, error) {

	var list []*models.AdminLog
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.AdminName) {
		orm = orm.Where("admin_name", request.AdminName)
	}
	if !gconv.IsEmpty(request.Data) {
		orm = orm.Where("data", request.Data)
	}
	if !gconv.IsEmpty(request.Method) {
		orm = orm.Where("method", request.Method)
	}
	if !gconv.IsEmpty(request.Path) {
		orm = orm.Where("path", request.Path)
	}
	if !gconv.IsEmpty(request.RequestIp) {
		orm = orm.Where("request_ip", request.RequestIp)
	}
	if !gconv.IsEmpty(request.RouteDesc) {
		orm = orm.Where("route_desc", request.RouteDesc)
	}
	if !gconv.IsEmpty(request.RouteName) {
		orm = orm.Where("route_name", request.RouteName)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}

	if request.Page > 0 && request.PageSize > 0 {
		orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)
	} else {
		orm.Order("id desc").Get(&list)
		count = int64(len(list))
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *AdminLogService) GetAll(request requests.AdminLogRequest) ([]*models.AdminLog, error) {

	var list []*models.AdminLog

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.AdminName) {
		orm = orm.Where("admin_name", request.AdminName)
	}
	if !gconv.IsEmpty(request.Data) {
		orm = orm.Where("data", request.Data)
	}
	if !gconv.IsEmpty(request.Method) {
		orm = orm.Where("method", request.Method)
	}
	if !gconv.IsEmpty(request.Path) {
		orm = orm.Where("path", request.Path)
	}
	if !gconv.IsEmpty(request.RequestIp) {
		orm = orm.Where("request_ip", request.RequestIp)
	}
	if !gconv.IsEmpty(request.RouteDesc) {
		orm = orm.Where("route_desc", request.RouteDesc)
	}
	if !gconv.IsEmpty(request.RouteName) {
		orm = orm.Where("route_name", request.RouteName)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminLogService) GetOne(id int64) (res models.AdminLog, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var adminLog models.AdminLog
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&adminLog)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return adminLog, nil
}

func (r *AdminLogService) Add(request requests.AdminLogRequest) (bool, error) {

	var adminLog models.AdminLog

	if !gconv.IsEmpty(request.AdminId) {
		adminLog.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.AdminName) {
		adminLog.AdminName = html.EscapeString(request.AdminName)
	}
	if !gconv.IsEmpty(request.Data) {
		adminLog.Data = html.EscapeString(request.Data)
	}
	if !gconv.IsEmpty(request.Method) {
		adminLog.Method = html.EscapeString(request.Method)
	}
	if !gconv.IsEmpty(request.Path) {
		adminLog.Path = html.EscapeString(request.Path)
	}
	if !gconv.IsEmpty(request.RequestIp) {
		adminLog.RequestIp = html.EscapeString(request.RequestIp)
	}
	if !gconv.IsEmpty(request.RouteDesc) {
		adminLog.RouteDesc = html.EscapeString(request.RouteDesc)
	}
	if !gconv.IsEmpty(request.RouteName) {
		adminLog.RouteName = html.EscapeString(request.RouteName)
	}
	if !gconv.IsEmpty(request.Url) {
		adminLog.Url = html.EscapeString(request.Url)
	}

	err := facades.Orm().Query().Create(&adminLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminLogService) Save(request requests.AdminLogRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var adminLog models.AdminLog
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&adminLog)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.AdminId) {
		adminLog.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.AdminName) {
		adminLog.AdminName = html.EscapeString(request.AdminName)
	}
	if !gconv.IsEmpty(request.Data) {
		adminLog.Data = html.EscapeString(request.Data)
	}
	if !gconv.IsEmpty(request.Method) {
		adminLog.Method = html.EscapeString(request.Method)
	}
	if !gconv.IsEmpty(request.Path) {
		adminLog.Path = html.EscapeString(request.Path)
	}
	if !gconv.IsEmpty(request.RequestIp) {
		adminLog.RequestIp = html.EscapeString(request.RequestIp)
	}
	if !gconv.IsEmpty(request.RouteDesc) {
		adminLog.RouteDesc = html.EscapeString(request.RouteDesc)
	}
	if !gconv.IsEmpty(request.RouteName) {
		adminLog.RouteName = html.EscapeString(request.RouteName)
	}
	if !gconv.IsEmpty(request.Url) {
		adminLog.Url = html.EscapeString(request.Url)
	}

	err = facades.Orm().Query().Save(&adminLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminLogService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var adminLog models.AdminLog
	_, err := facades.Orm().Query().Where("id", id).Delete(&adminLog)
	if err != nil {
		return false, err
	}
	return true, nil
}
