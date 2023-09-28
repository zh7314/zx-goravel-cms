package admin

import (
	"errors"
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

	var list []*models.RequestLog
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Data) {
		orm = orm.Where("data", request.Data)
	}
	if !gconv.IsEmpty(request.Header) {
		orm = orm.Where("header", request.Header)
	}
	if !gconv.IsEmpty(request.Ip) {
		orm = orm.Where("ip", request.Ip)
	}
	if !gconv.IsEmpty(request.Method) {
		orm = orm.Where("method", request.Method)
	}
	if !gconv.IsEmpty(request.Params) {
		orm = orm.Where("params", request.Params)
	}
	if !request.ReturnAt.IsZero() {
		orm = orm.Where("return_at", request.ReturnAt)
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

func (r *RequestLogService) GetAll(request requests.RequestLogRequest) ([]*models.RequestLog, error) {

	var list []*models.RequestLog

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Data) {
		orm = orm.Where("data", request.Data)
	}
	if !gconv.IsEmpty(request.Header) {
		orm = orm.Where("header", request.Header)
	}
	if !gconv.IsEmpty(request.Ip) {
		orm = orm.Where("ip", request.Ip)
	}
	if !gconv.IsEmpty(request.Method) {
		orm = orm.Where("method", request.Method)
	}
	if !gconv.IsEmpty(request.Params) {
		orm = orm.Where("params", request.Params)
	}
	if !request.ReturnAt.IsZero() {
		orm = orm.Where("return_at", request.ReturnAt)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *RequestLogService) GetOne(id int64) (res models.RequestLog, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var requestLog models.RequestLog
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&requestLog)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return requestLog, nil
}

func (r *RequestLogService) Add(request requests.RequestLogRequest) (bool, error) {

	var requestLog models.RequestLog

	if !gconv.IsEmpty(request.Data) {
		requestLog.Data = html.EscapeString(request.Data)
	}
	if !gconv.IsEmpty(request.Header) {
		requestLog.Header = html.EscapeString(request.Header)
	}
	if !gconv.IsEmpty(request.Ip) {
		requestLog.Ip = html.EscapeString(request.Ip)
	}
	if !gconv.IsEmpty(request.Method) {
		requestLog.Method = html.EscapeString(request.Method)
	}
	if !gconv.IsEmpty(request.Params) {
		requestLog.Params = html.EscapeString(request.Params)
	}
	if !request.ReturnAt.IsZero() {
		requestLog.ReturnAt = &request.ReturnAt
	}
	if !gconv.IsEmpty(request.Url) {
		requestLog.Url = html.EscapeString(request.Url)
	}

	err := facades.Orm().Query().Create(&requestLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RequestLogService) Save(request requests.RequestLogRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var requestLog models.RequestLog
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&requestLog)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Data) {
		requestLog.Data = html.EscapeString(request.Data)
	}
	if !gconv.IsEmpty(request.Header) {
		requestLog.Header = html.EscapeString(request.Header)
	}
	if !gconv.IsEmpty(request.Ip) {
		requestLog.Ip = html.EscapeString(request.Ip)
	}
	if !gconv.IsEmpty(request.Method) {
		requestLog.Method = html.EscapeString(request.Method)
	}
	if !gconv.IsEmpty(request.Params) {
		requestLog.Params = html.EscapeString(request.Params)
	}
	if !request.ReturnAt.IsZero() {
		requestLog.ReturnAt = &request.ReturnAt
	}
	if !gconv.IsEmpty(request.Url) {
		requestLog.Url = html.EscapeString(request.Url)
	}

	err = facades.Orm().Query().Save(&requestLog)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RequestLogService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var requestLog models.RequestLog
	_, err := facades.Orm().Query().Where("id", id).Delete(&requestLog)
	if err != nil {
		return false, err
	}
	return true, nil
}
