package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type ProductService struct {
}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (r *ProductService) GetList(request requests.ProductRequest) (map[string]interface{}, error) {

	var list []models.Product
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.ProductCateId) {
	orm.Where("product_cate_id", request.ProductCateId)
}
if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.ShortTitle) {
	orm.Where("short_title", request.ShortTitle)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.ViewCount) {
	orm.Where("view_count", request.ViewCount)
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
if !gconv.IsEmpty(request.StartTime) {
	orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.EndTime) {
	orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Pics) {
	orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.VideoUrl) {
	orm.Where("video_url", request.VideoUrl)
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

func (r *ProductService) GetAll(request requests.ProductRequest) ([]models.Product, error) {

	var list []models.Product

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.ProductCateId) {
	orm.Where("product_cate_id", request.ProductCateId)
}
if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.ShortTitle) {
	orm.Where("short_title", request.ShortTitle)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.ViewCount) {
	orm.Where("view_count", request.ViewCount)
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
if !gconv.IsEmpty(request.StartTime) {
	orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.EndTime) {
	orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Pics) {
	orm.Where("pics", request.Pics)
}
if !gconv.IsEmpty(request.VideoUrl) {
	orm.Where("video_url", request.VideoUrl)
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

func (r *ProductService) GetOne(id int64) (res models.Product, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var product models.Product
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&product)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return product, nil
}

func (r *ProductService) Add(request requests.ProductRequest) (bool, error) {

	var product models.Product

	product.ProductCateId = request.ProductCateId
product.Title = html.EscapeString(request.Title)
product.ShortTitle = html.EscapeString(request.ShortTitle)
product.Content = html.EscapeString(request.Content)
product.AdminId = request.AdminId
product.ViewCount = request.ViewCount
product.IsShow = request.IsShow
product.Sort = request.Sort
product.Url = html.EscapeString(request.Url)
product.StartTime = request.StartTime
product.EndTime = request.EndTime
product.Pic = html.EscapeString(request.Pic)
product.Pics = html.EscapeString(request.Pics)
product.VideoUrl = html.EscapeString(request.VideoUrl)
product.Platform = html.EscapeString(request.Platform)
product.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&product)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *ProductService) Save(request requests.ProductRequest) (bool, error) {

	var product models.Product

	product.ID = request.ID
	product.ProductCateId = request.ProductCateId
product.Title = html.EscapeString(request.Title)
product.ShortTitle = html.EscapeString(request.ShortTitle)
product.Content = html.EscapeString(request.Content)
product.AdminId = request.AdminId
product.ViewCount = request.ViewCount
product.IsShow = request.IsShow
product.Sort = request.Sort
product.Url = html.EscapeString(request.Url)
product.StartTime = request.StartTime
product.EndTime = request.EndTime
product.Pic = html.EscapeString(request.Pic)
product.Pics = html.EscapeString(request.Pics)
product.VideoUrl = html.EscapeString(request.VideoUrl)
product.Platform = html.EscapeString(request.Platform)
product.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&product)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProductService) Delete(id int64) (bool, error) {

	var admin models.Product
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
