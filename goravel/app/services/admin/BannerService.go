package admin

import (
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

	var list []models.Banner
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.BannerCateId) {
	orm.Where("banner_cate_id", request.BannerCateId)
}
if !gconv.IsEmpty(request.EndTime) {
	orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.PicPath) {
	orm.Where("pic_path", request.PicPath)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.StartTime) {
	orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.VideoPath) {
	orm.Where("video_path", request.VideoPath)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *BannerService) GetAll(request requests.BannerRequest) ([]models.Banner, error) {

	var list []models.Banner

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.BannerCateId) {
	orm.Where("banner_cate_id", request.BannerCateId)
}
if !gconv.IsEmpty(request.EndTime) {
	orm.Where("end_time", request.EndTime)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.PicPath) {
	orm.Where("pic_path", request.PicPath)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.StartTime) {
	orm.Where("start_time", request.StartTime)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.VideoPath) {
	orm.Where("video_path", request.VideoPath)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *BannerService) Add(request requests.BannerRequest) (bool, error) {

	var banner models.Banner

	banner.AdminId = request.AdminId
banner.BannerCateId = request.BannerCateId
banner.EndTime = request.EndTime
banner.Lang = html.EscapeString(request.Lang)
banner.Name = html.EscapeString(request.Name)
banner.PicPath = html.EscapeString(request.PicPath)
banner.Platform = html.EscapeString(request.Platform)
banner.Sort = request.Sort
banner.StartTime = request.StartTime
banner.Url = html.EscapeString(request.Url)
banner.VideoPath = html.EscapeString(request.VideoPath)


	err := facades.Orm().Query().Create(&banner)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *BannerService) Save(request requests.BannerRequest) (bool, error) {

	var banner models.Banner

	banner.ID = request.ID
	banner.AdminId = request.AdminId
banner.BannerCateId = request.BannerCateId
banner.EndTime = request.EndTime
banner.Lang = html.EscapeString(request.Lang)
banner.Name = html.EscapeString(request.Name)
banner.PicPath = html.EscapeString(request.PicPath)
banner.Platform = html.EscapeString(request.Platform)
banner.Sort = request.Sort
banner.StartTime = request.StartTime
banner.Url = html.EscapeString(request.Url)
banner.VideoPath = html.EscapeString(request.VideoPath)


	err := facades.Orm().Query().Save(&banner)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *BannerService) Delete(id int64) (bool, error) {

	var admin models.Banner
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
