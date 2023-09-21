package admin

import (
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

	var list []models.ProductCate
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
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.Description) {
	orm.Where("description", request.Description)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
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

func (r *ProductCateService) GetAll(request requests.ProductCateRequest) ([]models.ProductCate, error) {

	var list []models.ProductCate

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
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.Description) {
	orm.Where("description", request.Description)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
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

func (r *ProductCateService) Add(request requests.ProductCateRequest) (bool, error) {

	var productCate models.ProductCate

	productCate.ParentId = request.ParentId
productCate.Name = html.EscapeString(request.Name)
productCate.IsShow = request.IsShow
productCate.Sort = request.Sort
productCate.Url = html.EscapeString(request.Url)
productCate.Description = html.EscapeString(request.Description)
productCate.Pic = html.EscapeString(request.Pic)
productCate.Platform = html.EscapeString(request.Platform)
productCate.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&productCate)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *ProductCateService) Save(request requests.ProductCateRequest) (bool, error) {

	var productCate models.ProductCate

	productCate.ID = request.ID
	productCate.ParentId = request.ParentId
productCate.Name = html.EscapeString(request.Name)
productCate.IsShow = request.IsShow
productCate.Sort = request.Sort
productCate.Url = html.EscapeString(request.Url)
productCate.Description = html.EscapeString(request.Description)
productCate.Pic = html.EscapeString(request.Pic)
productCate.Platform = html.EscapeString(request.Platform)
productCate.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&productCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProductCateService) Delete(id int64) (bool, error) {

	var admin models.ProductCate
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
