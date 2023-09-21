package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type JobOffersService struct {
}

func NewJobOffersService() *JobOffersService {
	return &JobOffersService{}
}

func (r *JobOffersService) GetList(request requests.JobOffersRequest) (map[string]interface{}, error) {

	var list []models.JobOffers
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.SalaryRange) {
	orm.Where("salary_range", request.SalaryRange)
}
if !gconv.IsEmpty(request.Place) {
	orm.Where("place", request.Place)
}
if !gconv.IsEmpty(request.Number) {
	orm.Where("number", request.Number)
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

func (r *JobOffersService) GetAll(request requests.JobOffersRequest) ([]models.JobOffers, error) {

	var list []models.JobOffers

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.SalaryRange) {
	orm.Where("salary_range", request.SalaryRange)
}
if !gconv.IsEmpty(request.Place) {
	orm.Where("place", request.Place)
}
if !gconv.IsEmpty(request.Number) {
	orm.Where("number", request.Number)
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

func (r *JobOffersService) Add(request requests.JobOffersRequest) (bool, error) {

	var jobOffers models.JobOffers

	jobOffers.Title = html.EscapeString(request.Title)
jobOffers.Url = html.EscapeString(request.Url)
jobOffers.Content = html.EscapeString(request.Content)
jobOffers.IsShow = request.IsShow
jobOffers.Sort = request.Sort
jobOffers.SalaryRange = html.EscapeString(request.SalaryRange)
jobOffers.Place = html.EscapeString(request.Place)
jobOffers.Number = html.EscapeString(request.Number)
jobOffers.Platform = html.EscapeString(request.Platform)
jobOffers.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&jobOffers)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *JobOffersService) Save(request requests.JobOffersRequest) (bool, error) {

	var jobOffers models.JobOffers

	jobOffers.ID = request.ID
	jobOffers.Title = html.EscapeString(request.Title)
jobOffers.Url = html.EscapeString(request.Url)
jobOffers.Content = html.EscapeString(request.Content)
jobOffers.IsShow = request.IsShow
jobOffers.Sort = request.Sort
jobOffers.SalaryRange = html.EscapeString(request.SalaryRange)
jobOffers.Place = html.EscapeString(request.Place)
jobOffers.Number = html.EscapeString(request.Number)
jobOffers.Platform = html.EscapeString(request.Platform)
jobOffers.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&jobOffers)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *JobOffersService) Delete(id int64) (bool, error) {

	var admin models.JobOffers
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
