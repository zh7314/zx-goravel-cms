package admin

import (
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

	var list []models.Platform
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Value) {
	orm.Where("value", request.Value)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *PlatformService) GetAll(request requests.PlatformRequest) ([]models.Platform, error) {

	var list []models.Platform

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Value) {
	orm.Where("value", request.Value)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *PlatformService) Add(request requests.PlatformRequest) (bool, error) {

	var platform models.Platform

	platform.Name = html.EscapeString(request.Name)
platform.Sort = request.Sort
platform.Value = html.EscapeString(request.Value)


	err := facades.Orm().Query().Create(&platform)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *PlatformService) Save(request requests.PlatformRequest) (bool, error) {

	var platform models.Platform

	platform.ID = request.ID
	platform.Name = html.EscapeString(request.Name)
platform.Sort = request.Sort
platform.Value = html.EscapeString(request.Value)


	err := facades.Orm().Query().Save(&platform)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *PlatformService) Delete(id int64) (bool, error) {

	var admin models.Platform
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
