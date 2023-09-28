package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type BannerService struct {
}

func NewBannerService() *BannerService {
	return &BannerService{}
}

func (r *BannerService) GetList(request requests.BannerRequest) (map[string]interface{}, error) {

	var list []*models.Banner
	var count int64

	orm := facades.Orm().Query().With("BannerCate")

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.BannerCateId) {
		orm = orm.Where("banner_cate_id", request.BannerCateId)
	}
	if !request.EndTime.IsZero() {
		orm = orm.Where("end_time", request.EndTime)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.PicPath) {
		orm = orm.Where("pic_path", request.PicPath)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !request.StartTime.IsZero() {
		orm = orm.Where("start_time", request.StartTime)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}
	if !gconv.IsEmpty(request.VideoPath) {
		orm = orm.Where("video_path", request.VideoPath)
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

func (r *BannerService) GetAll(request requests.BannerRequest) ([]*models.Banner, error) {

	var list []*models.Banner

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.BannerCateId) {
		orm = orm.Where("banner_cate_id", request.BannerCateId)
	}
	if !request.EndTime.IsZero() {
		orm = orm.Where("end_time", request.EndTime)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		orm = orm.Where("name", request.Name)
	}
	if !gconv.IsEmpty(request.PicPath) {
		orm = orm.Where("pic_path", request.PicPath)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !request.StartTime.IsZero() {
		orm = orm.Where("start_time", request.StartTime)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}
	if !gconv.IsEmpty(request.VideoPath) {
		orm = orm.Where("video_path", request.VideoPath)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *BannerService) GetOne(id int64) (res models.Banner, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var banner models.Banner
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&banner)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return banner, nil
}

func (r *BannerService) Add(request requests.BannerRequest) (bool, error) {

	var banner models.Banner

	if !gconv.IsEmpty(request.AdminId) {
		banner.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.BannerCateId) {
		banner.BannerCateId = request.BannerCateId
	}
	if !request.EndTime.IsZero() {
		banner.EndTime = &request.EndTime
	}
	if !gconv.IsEmpty(request.Lang) {
		banner.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		banner.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.PicPath) {
		banner.PicPath = html.EscapeString(request.PicPath)
	}
	if !gconv.IsEmpty(request.Platform) {
		banner.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		banner.Sort = request.Sort
	}
	if !request.StartTime.IsZero() {
		banner.StartTime = &request.StartTime
	}
	if !gconv.IsEmpty(request.Url) {
		banner.Url = html.EscapeString(request.Url)
	}
	if !gconv.IsEmpty(request.VideoPath) {
		banner.VideoPath = html.EscapeString(request.VideoPath)
	}

	err := facades.Orm().Query().Create(&banner)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BannerService) Save(request requests.BannerRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var banner models.Banner
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&banner)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.AdminId) {
		banner.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.BannerCateId) {
		banner.BannerCateId = request.BannerCateId
	}
	if !request.EndTime.IsZero() {
		banner.EndTime = &request.EndTime
	}
	if !gconv.IsEmpty(request.Lang) {
		banner.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		banner.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.PicPath) {
		banner.PicPath = html.EscapeString(request.PicPath)
	}
	if !gconv.IsEmpty(request.Platform) {
		banner.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		banner.Sort = request.Sort
	}
	if !request.StartTime.IsZero() {
		banner.StartTime = &request.StartTime
	}
	if !gconv.IsEmpty(request.Url) {
		banner.Url = html.EscapeString(request.Url)
	}
	if !gconv.IsEmpty(request.VideoPath) {
		banner.VideoPath = html.EscapeString(request.VideoPath)
	}

	err = facades.Orm().Query().Save(&banner)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BannerService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var banner models.Banner
	_, err := facades.Orm().Query().Where("id", id).Delete(&banner)
	if err != nil {
		return false, err
	}
	return true, nil
}
