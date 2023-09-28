package admin

import (
	"errors"
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

	var list []*models.JobOffers
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Number) {
		orm = orm.Where("number", request.Number)
	}
	if !gconv.IsEmpty(request.Place) {
		orm = orm.Where("place", request.Place)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.SalaryRange) {
		orm = orm.Where("salary_range", request.SalaryRange)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}

	if request.Page > 0 && request.PageSize > 0 {
		orm.Order("sort asc").Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)
	} else {
		orm.Order("sort asc").Order("id desc").Get(&list)
		count = int64(len(list))
	}

	if len(list) > 0 {
		for _, v := range list {
			v.Content = html.UnescapeString(v.Content)
		}
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *JobOffersService) GetAll(request requests.JobOffersRequest) ([]*models.JobOffers, error) {

	var list []*models.JobOffers

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Number) {
		orm = orm.Where("number", request.Number)
	}
	if !gconv.IsEmpty(request.Place) {
		orm = orm.Where("place", request.Place)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.SalaryRange) {
		orm = orm.Where("salary_range", request.SalaryRange)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *JobOffersService) GetOne(id int64) (res models.JobOffers, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var jobOffers models.JobOffers
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&jobOffers)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return jobOffers, nil
}

func (r *JobOffersService) Add(request requests.JobOffersRequest) (bool, error) {

	var jobOffers models.JobOffers

	if !gconv.IsEmpty(request.Content) {
		jobOffers.Content = html.EscapeString(request.Content)
	}
	if !gconv.IsEmpty(request.IsShow) {
		jobOffers.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		jobOffers.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Number) {
		jobOffers.Number = html.EscapeString(request.Number)
	}
	if !gconv.IsEmpty(request.Place) {
		jobOffers.Place = html.EscapeString(request.Place)
	}
	if !gconv.IsEmpty(request.Platform) {
		jobOffers.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.SalaryRange) {
		jobOffers.SalaryRange = html.EscapeString(request.SalaryRange)
	}
	if !gconv.IsEmpty(request.Sort) {
		jobOffers.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Title) {
		jobOffers.Title = html.EscapeString(request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		jobOffers.Url = html.EscapeString(request.Url)
	}

	err := facades.Orm().Query().Create(&jobOffers)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *JobOffersService) Save(request requests.JobOffersRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var jobOffers models.JobOffers
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&jobOffers)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Content) {
		jobOffers.Content = html.EscapeString(request.Content)
	}
	if !gconv.IsEmpty(request.IsShow) {
		jobOffers.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		jobOffers.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Number) {
		jobOffers.Number = html.EscapeString(request.Number)
	}
	if !gconv.IsEmpty(request.Place) {
		jobOffers.Place = html.EscapeString(request.Place)
	}
	if !gconv.IsEmpty(request.Platform) {
		jobOffers.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.SalaryRange) {
		jobOffers.SalaryRange = html.EscapeString(request.SalaryRange)
	}
	if !gconv.IsEmpty(request.Sort) {
		jobOffers.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Title) {
		jobOffers.Title = html.EscapeString(request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		jobOffers.Url = html.EscapeString(request.Url)
	}

	err = facades.Orm().Query().Save(&jobOffers)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *JobOffersService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var jobOffers models.JobOffers
	_, err := facades.Orm().Query().Where("id", id).Delete(&jobOffers)
	if err != nil {
		return false, err
	}
	return true, nil
}
