package admin

import (
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

	if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *DownloadCateService) GetAll(request requests.DownloadCateRequest) ([]models.DownloadCate, error) {

	var list []models.DownloadCate

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.ParentId) {
	orm.Where("parent_id", request.ParentId)
}
if !gconv.IsEmpty(request.Pic) {
	orm.Where("pic", request.Pic)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *DownloadCateService) Add(request requests.DownloadCateRequest) (bool, error) {

	var downloadCate models.DownloadCate

	downloadCate.IsShow = request.IsShow
downloadCate.Lang = html.EscapeString(request.Lang)
downloadCate.Name = html.EscapeString(request.Name)
downloadCate.ParentId = request.ParentId
downloadCate.Pic = html.EscapeString(request.Pic)
downloadCate.Platform = html.EscapeString(request.Platform)
downloadCate.Sort = request.Sort


	err := facades.Orm().Query().Create(&downloadCate)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *DownloadCateService) Save(request requests.DownloadCateRequest) (bool, error) {

	var downloadCate models.DownloadCate

	downloadCate.ID = request.ID
	downloadCate.IsShow = request.IsShow
downloadCate.Lang = html.EscapeString(request.Lang)
downloadCate.Name = html.EscapeString(request.Name)
downloadCate.ParentId = request.ParentId
downloadCate.Pic = html.EscapeString(request.Pic)
downloadCate.Platform = html.EscapeString(request.Platform)
downloadCate.Sort = request.Sort


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
