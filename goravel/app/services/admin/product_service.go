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

	var list []*models.Product
	var count int64

	orm := facades.Orm().Query().With("ProductCate")

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !request.EndTime.IsZero() {
		orm = orm.Where("end_time", request.EndTime)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		orm = orm.Where("pics", request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.ProductCateId) {
		orm = orm.Where("product_cate_id", request.ProductCateId)
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		orm = orm.Where("short_title", request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !request.StartTime.IsZero() {
		orm = orm.Where("start_time", request.StartTime)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}
	if !gconv.IsEmpty(request.VideoUrl) {
		orm = orm.Where("video_url", request.VideoUrl)
	}
	if !gconv.IsEmpty(request.ViewCount) {
		orm = orm.Where("view_count", request.ViewCount)
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

func (r *ProductService) GetAll(request requests.ProductRequest) ([]*models.Product, error) {

	var list []*models.Product

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !request.EndTime.IsZero() {
		orm = orm.Where("end_time", request.EndTime)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		orm = orm.Where("pics", request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.ProductCateId) {
		orm = orm.Where("product_cate_id", request.ProductCateId)
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		orm = orm.Where("short_title", request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !request.StartTime.IsZero() {
		orm = orm.Where("start_time", request.StartTime)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}
	if !gconv.IsEmpty(request.VideoUrl) {
		orm = orm.Where("video_url", request.VideoUrl)
	}
	if !gconv.IsEmpty(request.ViewCount) {
		orm = orm.Where("view_count", request.ViewCount)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

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

	if !gconv.IsEmpty(request.AdminId) {
		product.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.Content) {
		product.Content = html.EscapeString(request.Content)
	}
	if !request.EndTime.IsZero() {
		product.EndTime = &request.EndTime
	}
	if !gconv.IsEmpty(request.IsShow) {
		product.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		product.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		product.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		product.Pics = html.EscapeString(request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		product.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.ProductCateId) {
		product.ProductCateId = request.ProductCateId
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		product.ShortTitle = html.EscapeString(request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		product.Sort = request.Sort
	}
	if !request.StartTime.IsZero() {
		product.StartTime = &request.StartTime
	}
	if !gconv.IsEmpty(request.Title) {
		product.Title = html.EscapeString(request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		product.Url = html.EscapeString(request.Url)
	}
	if !gconv.IsEmpty(request.VideoUrl) {
		product.VideoUrl = html.EscapeString(request.VideoUrl)
	}
	if !gconv.IsEmpty(request.ViewCount) {
		product.ViewCount = request.ViewCount
	}

	err := facades.Orm().Query().Create(&product)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProductService) Save(request requests.ProductRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var product models.Product
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&product)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.AdminId) {
		product.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.Content) {
		product.Content = html.EscapeString(request.Content)
	}
	if !request.EndTime.IsZero() {
		product.EndTime = &request.EndTime
	}
	if !gconv.IsEmpty(request.IsShow) {
		product.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		product.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		product.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		product.Pics = html.EscapeString(request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		product.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.ProductCateId) {
		product.ProductCateId = request.ProductCateId
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		product.ShortTitle = html.EscapeString(request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		product.Sort = request.Sort
	}
	if !request.StartTime.IsZero() {
		product.StartTime = &request.StartTime
	}
	if !gconv.IsEmpty(request.Title) {
		product.Title = html.EscapeString(request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		product.Url = html.EscapeString(request.Url)
	}
	if !gconv.IsEmpty(request.VideoUrl) {
		product.VideoUrl = html.EscapeString(request.VideoUrl)
	}
	if !gconv.IsEmpty(request.ViewCount) {
		product.ViewCount = request.ViewCount
	}

	err = facades.Orm().Query().Save(&product)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ProductService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var product models.Product
	_, err := facades.Orm().Query().Where("id", id).Delete(&product)
	if err != nil {
		return false, err
	}
	return true, nil
}
