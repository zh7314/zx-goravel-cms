package admin

import (
	"errors"
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

	var list []*models.Video
	var count int64

	orm := facades.Orm().Query().With("VideoCate")

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.Count) {
		orm = orm.Where("count", request.Count)
	}
	if !gconv.IsEmpty(request.File) {
		orm = orm.Where("file", request.File)
	}
	if !gconv.IsEmpty(request.Introduce) {
		orm = orm.Where("introduce", request.Introduce)
	}
	if !gconv.IsEmpty(request.IsLocal) {
		orm = orm.Where("is_local", request.IsLocal)
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
	if !gconv.IsEmpty(request.VideoCateId) {
		orm = orm.Where("video_cate_id", request.VideoCateId)
	}

	if request.Page > 0 && request.PageSize > 0 {
		orm.Order("sort asc").Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)
	} else {
		orm.Order("sort asc").Order("id desc").Get(&list)
		count = int64(len(list))
	}

	if len(list) > 0 {
		for _, v := range list {
			v.Introduce = html.UnescapeString(v.Introduce)
		}
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *VideoService) GetAll(request requests.VideoRequest) ([]*models.Video, error) {

	var list []*models.Video

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.Count) {
		orm = orm.Where("count", request.Count)
	}
	if !gconv.IsEmpty(request.File) {
		orm = orm.Where("file", request.File)
	}
	if !gconv.IsEmpty(request.Introduce) {
		orm = orm.Where("introduce", request.Introduce)
	}
	if !gconv.IsEmpty(request.IsLocal) {
		orm = orm.Where("is_local", request.IsLocal)
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
	if !gconv.IsEmpty(request.VideoCateId) {
		orm = orm.Where("video_cate_id", request.VideoCateId)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *VideoService) GetOne(id int64) (res models.Video, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var video models.Video
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&video)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return video, nil
}

func (r *VideoService) Add(request requests.VideoRequest) (bool, error) {

	var video models.Video

	if !gconv.IsEmpty(request.AdminId) {
		video.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.Count) {
		video.Count = request.Count
	}
	if !gconv.IsEmpty(request.File) {
		video.File = html.EscapeString(request.File)
	}
	if !gconv.IsEmpty(request.Introduce) {
		video.Introduce = html.EscapeString(request.Introduce)
	}
	if !gconv.IsEmpty(request.IsLocal) {
		video.IsLocal = request.IsLocal
	}
	if !gconv.IsEmpty(request.IsShow) {
		video.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		video.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		video.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Pic) {
		video.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		video.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		video.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Url) {
		video.Url = html.EscapeString(request.Url)
	}
	if !gconv.IsEmpty(request.VideoCateId) {
		video.VideoCateId = request.VideoCateId
	}

	err := facades.Orm().Query().Create(&video)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *VideoService) Save(request requests.VideoRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var video models.Video
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&video)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.AdminId) {
		video.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.Count) {
		video.Count = request.Count
	}
	if !gconv.IsEmpty(request.File) {
		video.File = html.EscapeString(request.File)
	}
	if !gconv.IsEmpty(request.Introduce) {
		video.Introduce = html.EscapeString(request.Introduce)
	}
	if !gconv.IsEmpty(request.IsLocal) {
		video.IsLocal = request.IsLocal
	}
	if !gconv.IsEmpty(request.IsShow) {
		video.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		video.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Name) {
		video.Name = html.EscapeString(request.Name)
	}
	if !gconv.IsEmpty(request.Pic) {
		video.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		video.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		video.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Url) {
		video.Url = html.EscapeString(request.Url)
	}
	if !gconv.IsEmpty(request.VideoCateId) {
		video.VideoCateId = request.VideoCateId
	}

	err = facades.Orm().Query().Save(&video)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *VideoService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var video models.Video
	_, err := facades.Orm().Query().Where("id", id).Delete(&video)
	if err != nil {
		return false, err
	}
	return true, nil
}
