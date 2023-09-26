package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type VideoCateService struct {
}

func NewVideoCateService() *VideoCateService {
	return &VideoCateService{}
}

func (r *VideoCateService) GetList(request requests.VideoCateRequest) (map[string]interface{}, error) {

	var list []models.VideoCate
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

func (r *VideoCateService) GetAll(request requests.VideoCateRequest) ([]models.VideoCate, error) {

	var list []models.VideoCate

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
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *VideoCateService) GetOne(id int64) (res models.VideoCate, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var videoCate models.VideoCate
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&videoCate)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return videoCate, nil
}

func (r *VideoCateService) Add(request requests.VideoCateRequest) (bool, error) {

	var videoCate models.VideoCate

	videoCate.ParentId = request.ParentId
videoCate.Name = html.EscapeString(request.Name)
videoCate.IsShow = request.IsShow
videoCate.Sort = request.Sort
videoCate.Platform = html.EscapeString(request.Platform)
videoCate.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&videoCate)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *VideoCateService) Save(request requests.VideoCateRequest) (bool, error) {

	var videoCate models.VideoCate

	videoCate.ID = request.ID
	videoCate.ParentId = request.ParentId
videoCate.Name = html.EscapeString(request.Name)
videoCate.IsShow = request.IsShow
videoCate.Sort = request.Sort
videoCate.Platform = html.EscapeString(request.Platform)
videoCate.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&videoCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *VideoCateService) Delete(id int64) (bool, error) {

	var admin models.VideoCate
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
