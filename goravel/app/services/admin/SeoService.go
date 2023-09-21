package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type SeoService struct {
}

func NewSeoService() *SeoService {
	return &SeoService{}
}

func (r *SeoService) GetList(request requests.SeoRequest) (map[string]interface{}, error) {

	var list []models.Seo
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Keyword) {
	orm.Where("keyword", request.Keyword)
}
if !gconv.IsEmpty(request.Description) {
	orm.Where("description", request.Description)
}
if !gconv.IsEmpty(request.Position) {
	orm.Where("position", request.Position)
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

func (r *SeoService) GetAll(request requests.SeoRequest) ([]models.Seo, error) {

	var list []models.Seo

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Keyword) {
	orm.Where("keyword", request.Keyword)
}
if !gconv.IsEmpty(request.Description) {
	orm.Where("description", request.Description)
}
if !gconv.IsEmpty(request.Position) {
	orm.Where("position", request.Position)
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

func (r *SeoService) Add(request requests.SeoRequest) (bool, error) {

	var seo models.Seo

	seo.Title = html.EscapeString(request.Title)
seo.Keyword = html.EscapeString(request.Keyword)
seo.Description = html.EscapeString(request.Description)
seo.Position = html.EscapeString(request.Position)
seo.IsShow = request.IsShow
seo.Sort = request.Sort
seo.Platform = html.EscapeString(request.Platform)
seo.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&seo)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *SeoService) Save(request requests.SeoRequest) (bool, error) {

	var seo models.Seo

	seo.ID = request.ID
	seo.Title = html.EscapeString(request.Title)
seo.Keyword = html.EscapeString(request.Keyword)
seo.Description = html.EscapeString(request.Description)
seo.Position = html.EscapeString(request.Position)
seo.IsShow = request.IsShow
seo.Sort = request.Sort
seo.Platform = html.EscapeString(request.Platform)
seo.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&seo)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *SeoService) Delete(id int64) (bool, error) {

	var admin models.Seo
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
