package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type NewsCateService struct {
}

func NewNewsCateService() *NewsCateService {
	return &NewsCateService{}
}

func (r *NewsCateService) GetList(request requests.NewsCateRequest) (map[string]interface{}, error) {

	var list []models.NewsCate
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *NewsCateService) GetAll(request requests.NewsCateRequest) ([]models.NewsCate, error) {

	var list []models.NewsCate

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *NewsCateService) Add(request requests.NewsCateRequest) (bool, error) {

	var newsCate models.NewsCate

	newsCate.ParentId = request.ParentId
newsCate.Name = html.EscapeString(request.Name)
newsCate.IsShow = request.IsShow
newsCate.Sort = request.Sort
newsCate.Platform = html.EscapeString(request.Platform)
newsCate.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&newsCate)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *NewsCateService) Save(request requests.NewsCateRequest) (bool, error) {

	var newsCate models.NewsCate

	newsCate.ID = request.ID
	newsCate.ParentId = request.ParentId
newsCate.Name = html.EscapeString(request.Name)
newsCate.IsShow = request.IsShow
newsCate.Sort = request.Sort
newsCate.Platform = html.EscapeString(request.Platform)
newsCate.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&newsCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *NewsCateService) Delete(id int64) (bool, error) {

	var admin models.NewsCate
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
