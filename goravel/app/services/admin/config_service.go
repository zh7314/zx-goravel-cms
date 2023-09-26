package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type ConfigService struct {
}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (r *ConfigService) GetList(request requests.ConfigRequest) (map[string]interface{}, error) {

	var list []models.Config
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Type) {
	orm.Where("type", request.Type)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
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

func (r *ConfigService) GetAll(request requests.ConfigRequest) ([]models.Config, error) {

	var list []models.Config

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Type) {
	orm.Where("type", request.Type)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Value) {
	orm.Where("value", request.Value)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *ConfigService) GetOne(id int64) (res models.Config, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var config models.Config
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&config)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return config, nil
}

func (r *ConfigService) Add(request requests.ConfigRequest) (bool, error) {

	var config models.Config

	config.Type = html.EscapeString(request.Type)
config.Name = html.EscapeString(request.Name)
config.Value = html.EscapeString(request.Value)


	err := facades.Orm().Query().Create(&config)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *ConfigService) Save(request requests.ConfigRequest) (bool, error) {

	var config models.Config

	config.ID = request.ID
	config.Type = html.EscapeString(request.Type)
config.Name = html.EscapeString(request.Name)
config.Value = html.EscapeString(request.Value)


	err := facades.Orm().Query().Save(&config)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ConfigService) Delete(id int64) (bool, error) {

	var admin models.Config
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
