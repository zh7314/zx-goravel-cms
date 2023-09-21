package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type BannerCateService struct {
}

func NewBannerCateService() *BannerCateService {
	return &BannerCateService{}
}

func (r *BannerCateService) GetList(request requests.BannerCateRequest) (map[string]interface{}, error) {

	var list []models.BannerCate
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
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
if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *BannerCateService) GetAll(request requests.BannerCateRequest) ([]models.BannerCate, error) {

	var list []models.BannerCate

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
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
if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *BannerCateService) Add(request requests.BannerCateRequest) (bool, error) {

	var bannerCate models.BannerCate

	bannerCate.Name = html.EscapeString(request.Name)
bannerCate.IsShow = request.IsShow
bannerCate.Sort = request.Sort
bannerCate.Pic = html.EscapeString(request.Pic)
bannerCate.Platform = html.EscapeString(request.Platform)
bannerCate.Lang = html.EscapeString(request.Lang)
bannerCate.ParentId = request.ParentId


	err := facades.Orm().Query().Create(&bannerCate)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *BannerCateService) Save(request requests.BannerCateRequest) (bool, error) {

	var bannerCate models.BannerCate

	bannerCate.ID = request.ID
	bannerCate.Name = html.EscapeString(request.Name)
bannerCate.IsShow = request.IsShow
bannerCate.Sort = request.Sort
bannerCate.Pic = html.EscapeString(request.Pic)
bannerCate.Platform = html.EscapeString(request.Platform)
bannerCate.Lang = html.EscapeString(request.Lang)
bannerCate.ParentId = request.ParentId


	err := facades.Orm().Query().Save(&bannerCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BannerCateService) Delete(id int64) (bool, error) {

	var admin models.BannerCate
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
