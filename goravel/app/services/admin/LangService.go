package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type LangService struct {
}

func NewLangService() *LangService {
	return &LangService{}
}

func (r *LangService) GetList(request requests.LangRequest) (map[string]interface{}, error) {

	var list []models.Lang
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
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

func (r *LangService) GetAll(request requests.LangRequest) ([]models.Lang, error) {

	var list []models.Lang

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
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

func (r *LangService) Add(request requests.LangRequest) (bool, error) {

	var lang models.Lang

	lang.Sort = request.Sort
lang.Name = html.EscapeString(request.Name)
lang.Value = html.EscapeString(request.Value)


	err := facades.Orm().Query().Create(&lang)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *LangService) Save(request requests.LangRequest) (bool, error) {

	var lang models.Lang

	lang.ID = request.ID
	lang.Sort = request.Sort
lang.Name = html.EscapeString(request.Name)
lang.Value = html.EscapeString(request.Value)


	err := facades.Orm().Query().Save(&lang)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *LangService) Delete(id int64) (bool, error) {

	var admin models.Lang
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
