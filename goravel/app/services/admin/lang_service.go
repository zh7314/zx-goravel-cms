package admin

import (
	"errors"
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

	var list []*models.Lang
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

func (r *LangService) GetAll(request requests.LangRequest) ([]*models.Lang, error) {

	var list []*models.Lang

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

func (r *LangService) GetOne(id int64) (res models.Lang, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var lang models.Lang
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&lang)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return lang, nil
}

func (r *LangService) Add(request requests.LangRequest) (bool, error) {

	var lang models.Lang

	if !gconv.IsEmpty(request.Name) {
		lang.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Sort) {
		lang.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Value) {
		lang.Value = html.EscapeString(request.Value)
	}

	err := facades.Orm().Query().Create(&lang)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *LangService) Save(request requests.LangRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var lang models.Lang
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&lang)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Name) {
		lang.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Sort) {
		lang.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Value) {
		lang.Value = html.EscapeString(request.Value)
	}

	err = facades.Orm().Query().Save(&lang)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *LangService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var lang models.Lang
	_, err := facades.Orm().Query().Where("id", id).Delete(&lang)
	if err != nil {
		return false, err
	}
	return true, nil
}
