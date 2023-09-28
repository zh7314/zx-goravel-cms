package admin

import (
	"errors"
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

	var list []*models.NewsCate
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		orm = orm.Where("parent_id", request.ParentId)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
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

func (r *NewsCateService) GetAll(request requests.NewsCateRequest) ([]*models.NewsCate, error) {

	var list []*models.NewsCate

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		orm = orm.Where("parent_id", request.ParentId)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *NewsCateService) GetOne(id int64) (res models.NewsCate, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var newsCate models.NewsCate
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&newsCate)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return newsCate, nil
}

func (r *NewsCateService) Add(request requests.NewsCateRequest) (bool, error) {

	var newsCate models.NewsCate

	if !gconv.IsEmpty(request.IsShow) {
		newsCate.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		newsCate.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		newsCate.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		newsCate.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Platform) {
		newsCate.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		newsCate.Sort = request.Sort
	}

	err := facades.Orm().Query().Create(&newsCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *NewsCateService) Save(request requests.NewsCateRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var newsCate models.NewsCate
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&newsCate)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.IsShow) {
		newsCate.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		newsCate.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		newsCate.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		newsCate.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Platform) {
		newsCate.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		newsCate.Sort = request.Sort
	}

	err = facades.Orm().Query().Save(&newsCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *NewsCateService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var newsCate models.NewsCate
	_, err := facades.Orm().Query().Where("id", id).Delete(&newsCate)
	if err != nil {
		return false, err
	}
	return true, nil
}
