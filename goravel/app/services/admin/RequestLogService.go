package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type RequestLogService struct {
}

func NewRequestLogService() *RequestLogService {
	return &RequestLogService{}
}

func (r *RequestLogService) GetList(request requests.RequestLogRequest) (map[string]interface{}, error) {

	var list []models.RequestLog
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Data) {
	orm.Where("data", request.Data)
}
if !gconv.IsEmpty(request.Header) {
	orm.Where("header", request.Header)
}
if !gconv.IsEmpty(request.Ip) {
	orm.Where("ip", request.Ip)
}
if !gconv.IsEmpty(request.Method) {
	orm.Where("method", request.Method)
}
if !gconv.IsEmpty(request.Params) {
	orm.Where("params", request.Params)
}
if !gconv.IsEmpty(request.ReturnAt) {
	orm.Where("return_at", request.ReturnAt)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *RequestLogService) GetAll(request requests.RequestLogRequest) ([]models.RequestLog, error) {

	var list []models.RequestLog

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Data) {
	orm.Where("data", request.Data)
}
if !gconv.IsEmpty(request.Header) {
	orm.Where("header", request.Header)
}
if !gconv.IsEmpty(request.Ip) {
	orm.Where("ip", request.Ip)
}
if !gconv.IsEmpty(request.Method) {
	orm.Where("method", request.Method)
}
if !gconv.IsEmpty(request.Params) {
	orm.Where("params", request.Params)
}
if !gconv.IsEmpty(request.ReturnAt) {
	orm.Where("return_at", request.ReturnAt)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *RequestLogService) Add(request requests.RequestLogRequest) (bool, error) {

	var requestLog models.RequestLog

	requestLog.Data = html.EscapeString(request.Data)
requestLog.Header = html.EscapeString(request.Header)
requestLog.Ip = html.EscapeString(request.Ip)
requestLog.Method = html.EscapeString(request.Method)
requestLog.Params = html.EscapeString(request.Params)
requestLog.ReturnAt = request.ReturnAt
requestLog.Url = html.EscapeString(request.Url)


	err := facades.Orm().Query().Create(&requestLog)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *RequestLogService) Save(request requests.RequestLogRequest) (bool, error) {

	var requestLog models.RequestLog

	requestLog.ID = request.ID
	requestLog.Data = html.EscapeString(request.Data)
requestLog.Header = html.EscapeString(request.Header)
requestLog.Ip = html.EscapeString(request.Ip)
requestLog.Method = html.EscapeString(request.Method)
requestLog.Params = html.EscapeString(request.Params)
requestLog.ReturnAt = request.ReturnAt
requestLog.Url = html.EscapeString(request.Url)


	err := facades.Orm().Query().Save(&requestLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RequestLogService) Delete(id int64) (bool, error) {

	var admin models.RequestLog
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
