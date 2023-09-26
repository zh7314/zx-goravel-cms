package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type DownloadCateService struct {
}

func NewDownloadCateService() *DownloadCateService {
	return &DownloadCateService{}
}

func (r *DownloadCateService) GetList(request requests.DownloadCateRequest) (map[string]interface{}, error) {

	var list []models.DownloadCate
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
	orm = orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.IsShow) {
	orm = orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm = orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Pic) {
	orm = orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Platform) {
	orm = orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Lang) {
	orm = orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.ParentId) {
	orm = orm.Where("parent_id", request.ParentId)
}


	if request.Page > 0 && request.PageSize > 0 {
		orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)
	} else {
		orm.Order("id desc").Get(&list)
		count = int64(len(list))
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *DownloadCateService) GetAll(request requests.DownloadCateRequest) ([]models.DownloadCate, error) {

	var list []models.DownloadCate

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Name) {
	orm = orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.IsShow) {
	orm = orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm = orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Pic) {
	orm = orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Platform) {
	orm = orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Lang) {
	orm = orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.ParentId) {
	orm = orm.Where("parent_id", request.ParentId)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *DownloadCateService) GetOne(id int64) (res models.DownloadCate, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var downloadCate models.DownloadCate
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&downloadCate)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return downloadCate, nil
}

func (r *DownloadCateService) Add(request requests.DownloadCateRequest) (bool, error) {

	var downloadCate models.DownloadCate

	downloadCate.Name = html.EscapeString(request.Name)
downloadCate.IsShow = request.IsShow
downloadCate.Sort = request.Sort
downloadCate.Pic = html.EscapeString(request.Pic)
downloadCate.Platform = html.EscapeString(request.Platform)
downloadCate.Lang = html.EscapeString(request.Lang)
downloadCate.ParentId = request.ParentId


	err := facades.Orm().Query().Create(&downloadCate)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *DownloadCateService) Save(request requests.DownloadCateRequest) (bool, error) {

	var downloadCate models.DownloadCate

	downloadCate.ID = request.ID
	downloadCate.Name = html.EscapeString(request.Name)
downloadCate.IsShow = request.IsShow
downloadCate.Sort = request.Sort
downloadCate.Pic = html.EscapeString(request.Pic)
downloadCate.Platform = html.EscapeString(request.Platform)
downloadCate.Lang = html.EscapeString(request.Lang)
downloadCate.ParentId = request.ParentId


	err := facades.Orm().Query().Save(&downloadCate)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *DownloadCateService) Delete(id int64) (bool, error) {

	var admin models.DownloadCate
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
