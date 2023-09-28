package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type PlatformService struct {
}

func NewPlatformService() *PlatformService {
	return &PlatformService{}
}

func (r *PlatformService) GetList(request requests.PlatformRequest) (map[string]interface{}, error) {

	var list []*models.Platform
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Value) {
		orm = orm.Where("value", request.Value)
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

func (r *PlatformService) GetAll(request requests.PlatformRequest) ([]*models.Platform, error) {

	var list []*models.Platform

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Value) {
		orm = orm.Where("value", request.Value)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *PlatformService) GetOne(id int64) (res models.Platform, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var platform models.Platform
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&platform)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return platform, nil
}

func (r *PlatformService) Add(request requests.PlatformRequest) (bool, error) {

	var platform models.Platform

	if !gconv.IsEmpty(request.Name) {
		platform.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Sort) {
		platform.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Value) {
		platform.Value = html.EscapeString(request.Value)
	}

	err := facades.Orm().Query().Create(&platform)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *PlatformService) Save(request requests.PlatformRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var platform models.Platform
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&platform)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Name) {
		platform.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Sort) {
		platform.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Value) {
		platform.Value = html.EscapeString(request.Value)
	}

	err = facades.Orm().Query().Save(&platform)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *PlatformService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var platform models.Platform
	_, err := facades.Orm().Query().Where("id", id).Delete(&platform)
	if err != nil {
		return false, err
	}
	return true, nil
}
