package admin

import (
	"errors"
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

	var list []*models.BannerCate
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
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
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

func (r *BannerCateService) GetAll(request requests.BannerCateRequest) ([]*models.BannerCate, error) {

	var list []*models.BannerCate

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
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
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

func (r *BannerCateService) GetOne(id int64) (res models.BannerCate, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var bannerCate models.BannerCate
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&bannerCate)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return bannerCate, nil
}

func (r *BannerCateService) Add(request requests.BannerCateRequest) (bool, error) {

	var bannerCate models.BannerCate

	if !gconv.IsEmpty(request.IsShow) {
		bannerCate.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		bannerCate.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		bannerCate.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		bannerCate.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Pic) {
		bannerCate.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		bannerCate.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		bannerCate.Sort = request.Sort
	}

	err := facades.Orm().Query().Create(&bannerCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BannerCateService) Save(request requests.BannerCateRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var bannerCate models.BannerCate
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&bannerCate)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.IsShow) {
		bannerCate.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		bannerCate.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		bannerCate.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.ParentId) {
		bannerCate.ParentId = request.ParentId
	}
	if !gconv.IsEmpty(request.Pic) {
		bannerCate.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		bannerCate.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		bannerCate.Sort = request.Sort
	}

	err = facades.Orm().Query().Save(&bannerCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BannerCateService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var bannerCate models.BannerCate
	_, err := facades.Orm().Query().Where("id", id).Delete(&bannerCate)
	if err != nil {
		return false, err
	}
	return true, nil
}
