package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type ProductCateService struct {
}

func NewProductCateService() *ProductCateService {
	return &ProductCateService{}
}

func (r *ProductCateService) GetList(request requests.ProductCateRequest) (map[string]interface{}, error) {

	var list []*models.ProductCate
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Description) {
		orm = orm.Where("description", request.Description)
	}
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
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
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

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *ProductCateService) GetAll(request requests.ProductCateRequest) ([]*models.ProductCate, error) {

	var list []*models.ProductCate

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Description) {
		orm = orm.Where("description", request.Description)
	}
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
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *ProductCateService) GetOne(id int64) (res models.ProductCate, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var productCate models.ProductCate
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&productCate)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return productCate, nil
}

func (r *ProductCateService) Add(request requests.ProductCateRequest) (bool, error) {

	var productCate models.ProductCate

	if !gconv.IsEmpty(request.Description) {
		productCate.Description = html.EscapeString(request.Description)
	}
	if !gconv.IsEmpty(request.IsShow) {
		productCate.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		productCate.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		productCate.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		productCate.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Pic) {
		productCate.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		productCate.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		productCate.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Url) {
		productCate.Url = html.EscapeString(request.Url)
	}

	err := facades.Orm().Query().Create(&productCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProductCateService) Save(request requests.ProductCateRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var productCate models.ProductCate
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&productCate)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Description) {
		productCate.Description = html.EscapeString(request.Description)
	}
	if !gconv.IsEmpty(request.IsShow) {
		productCate.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		productCate.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		productCate.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		productCate.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Pic) {
		productCate.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		productCate.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		productCate.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Url) {
		productCate.Url = html.EscapeString(request.Url)
	}

	err = facades.Orm().Query().Save(&productCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProductCateService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var productCate models.ProductCate
	_, err := facades.Orm().Query().Where("id", id).Delete(&productCate)
	if err != nil {
		return false, err
	}
	return true, nil
}
