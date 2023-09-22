package admin

import (
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

	var list []models.AdminLog
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Method) {
		orm.Where("method", request.Method)
	}
	if !gconv.IsEmpty(request.Url) {
		orm.Where("url", request.Url)
	}
	if !gconv.IsEmpty(request.RouteName) {
		orm.Where("route_name", request.RouteName)
	}
	if !gconv.IsEmpty(request.Path) {
		orm.Where("path", request.Path)
	}
	if !gconv.IsEmpty(request.RequestIp) {
		orm.Where("request_ip", request.RequestIp)
	}
	if !gconv.IsEmpty(request.Data) {
		orm.Where("data", request.Data)
	}
	if !gconv.IsEmpty(request.AdminId) {
		orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.AdminName) {
		orm.Where("admin_name", request.AdminName)
	}
	if !gconv.IsEmpty(request.RouteDesc) {
		orm.Where("route_desc", request.RouteDesc)
	}

	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *AdminLogService) GetAll(request requests.AdminLogRequest) ([]models.AdminLog, error) {

	var list []models.AdminLog

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Method) {
		orm.Where("method", request.Method)
	}
	if !gconv.IsEmpty(request.Url) {
		orm.Where("url", request.Url)
	}
	if !gconv.IsEmpty(request.RouteName) {
		orm.Where("route_name", request.RouteName)
	}
	if !gconv.IsEmpty(request.Path) {
		orm.Where("path", request.Path)
	}
	if !gconv.IsEmpty(request.RequestIp) {
		orm.Where("request_ip", request.RequestIp)
	}
	if !gconv.IsEmpty(request.Data) {
		orm.Where("data", request.Data)
	}
	if !gconv.IsEmpty(request.AdminId) {
		orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.AdminName) {
		orm.Where("admin_name", request.AdminName)
	}
	if !gconv.IsEmpty(request.RouteDesc) {
		orm.Where("route_desc", request.RouteDesc)
	}

	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminLogService) Add(request requests.AdminLogRequest) (bool, error) {

	var adminLog models.AdminLog

	adminLog.Method = html.EscapeString(request.Method)
	adminLog.Url = html.EscapeString(request.Url)
	adminLog.RouteName = html.EscapeString(request.RouteName)
	adminLog.Path = html.EscapeString(request.Path)
	adminLog.RequestIp = html.EscapeString(request.RequestIp)
	adminLog.Data = html.EscapeString(request.Data)
	adminLog.AdminId = request.AdminId
	adminLog.AdminName = html.EscapeString(request.AdminName)
	adminLog.RouteDesc = html.EscapeString(request.RouteDesc)

	err := facades.Orm().Query().Create(&adminLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminLogService) Save(request requests.AdminLogRequest) (bool, error) {

	var adminLog models.AdminLog

	adminLog.ID = request.ID
	adminLog.Method = html.EscapeString(request.Method)
	adminLog.Url = html.EscapeString(request.Url)
	adminLog.RouteName = html.EscapeString(request.RouteName)
	adminLog.Path = html.EscapeString(request.Path)
	adminLog.RequestIp = html.EscapeString(request.RequestIp)
	adminLog.Data = html.EscapeString(request.Data)
	adminLog.AdminId = request.AdminId
	adminLog.AdminName = html.EscapeString(request.AdminName)
	adminLog.RouteDesc = html.EscapeString(request.RouteDesc)

	err := facades.Orm().Query().Save(&adminLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminLogService) Delete(id int64) (bool, error) {

	var admin models.AdminLog
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
