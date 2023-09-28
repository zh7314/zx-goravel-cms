package admin

import (
	"errors"
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

	var list []*models.Seo
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Description) {
		orm = orm.Where("description", request.Description)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Keyword) {
		orm = orm.Where("keyword", request.Keyword)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Position) {
		orm = orm.Where("position", request.Position)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
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

func (r *SeoService) GetAll(request requests.SeoRequest) ([]*models.Seo, error) {

	var list []*models.Seo

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Description) {
		orm = orm.Where("description", request.Description)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Keyword) {
		orm = orm.Where("keyword", request.Keyword)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Position) {
		orm = orm.Where("position", request.Position)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *SeoService) GetOne(id int64) (res models.Seo, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var seo models.Seo
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&seo)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return seo, nil
}

func (r *SeoService) Add(request requests.SeoRequest) (bool, error) {

	var seo models.Seo

	if !gconv.IsEmpty(request.Description) {
		seo.Description = html.EscapeString(request.Description)
	}
	if !gconv.IsEmpty(request.IsShow) {
		seo.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Keyword) {
		seo.Keyword = html.EscapeString(request.Keyword)
	}
	if !gconv.IsEmpty(request.Lang) {
		seo.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Platform) {
		seo.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Position) {
		seo.Position = html.EscapeString(request.Position)
	}
	if !gconv.IsEmpty(request.Sort) {
		seo.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Title) {
		seo.Title = html.EscapeString(request.Title)
	}

	err := facades.Orm().Query().Create(&seo)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *SeoService) Save(request requests.SeoRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var seo models.Seo
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&seo)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Description) {
		seo.Description = html.EscapeString(request.Description)
	}
	if !gconv.IsEmpty(request.IsShow) {
		seo.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Keyword) {
		seo.Keyword = html.EscapeString(request.Keyword)
	}
	if !gconv.IsEmpty(request.Lang) {
		seo.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Platform) {
		seo.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Position) {
		seo.Position = html.EscapeString(request.Position)
	}
	if !gconv.IsEmpty(request.Sort) {
		seo.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Title) {
		seo.Title = html.EscapeString(request.Title)
	}

	err = facades.Orm().Query().Save(&seo)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *SeoService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var seo models.Seo
	_, err := facades.Orm().Query().Where("id", id).Delete(&seo)
	if err != nil {
		return false, err
	}
	return true, nil
}
