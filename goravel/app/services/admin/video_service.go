package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type VideoService struct {
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (r *VideoService) GetList(request requests.VideoRequest) (map[string]interface{}, error) {

	var list []models.Video
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.VideoCateId) {
	orm.Where("video_cate_id", request.VideoCateId)
}
if !gconv.IsEmpty(request.IsLocal) {
	orm.Where("is_local", request.IsLocal)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Introduce) {
	orm.Where("introduce", request.Introduce)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.Count) {
	orm.Where("count", request.Count)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.File) {
	orm.Where("file", request.File)
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

func (r *VideoService) GetAll(request requests.VideoRequest) ([]models.Video, error) {

	var list []models.Video

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.VideoCateId) {
	orm.Where("video_cate_id", request.VideoCateId)
}
if !gconv.IsEmpty(request.IsLocal) {
	orm.Where("is_local", request.IsLocal)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Introduce) {
	orm.Where("introduce", request.Introduce)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.Count) {
	orm.Where("count", request.Count)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.File) {
	orm.Where("file", request.File)
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

func (r *VideoService) Add(request requests.VideoRequest) (bool, error) {

	var video models.Video

	video.VideoCateId = request.VideoCateId
video.IsLocal = request.IsLocal
video.Url = html.EscapeString(request.Url)
video.Name = html.EscapeString(request.Name)
video.Introduce = html.EscapeString(request.Introduce)
video.IsShow = request.IsShow
video.AdminId = request.AdminId
video.Count = request.Count
video.Sort = request.Sort
video.Pic = html.EscapeString(request.Pic)
video.File = html.EscapeString(request.File)
video.Platform = html.EscapeString(request.Platform)
video.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&video)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *VideoService) Save(request requests.VideoRequest) (bool, error) {

	var video models.Video

	video.ID = request.ID
	video.VideoCateId = request.VideoCateId
video.IsLocal = request.IsLocal
video.Url = html.EscapeString(request.Url)
video.Name = html.EscapeString(request.Name)
video.Introduce = html.EscapeString(request.Introduce)
video.IsShow = request.IsShow
video.AdminId = request.AdminId
video.Count = request.Count
video.Sort = request.Sort
video.Pic = html.EscapeString(request.Pic)
video.File = html.EscapeString(request.File)
video.Platform = html.EscapeString(request.Platform)
video.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&video)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *VideoService) Delete(id int64) (bool, error) {

	var admin models.Video
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
