package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type DownloadService struct {
}

func NewDownloadService() *DownloadService {
	return &DownloadService{}
}

func (r *DownloadService) GetList(request requests.DownloadRequest) (map[string]interface{}, error) {

	var list []models.Download
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Introduction) {
	orm.Where("introduction", request.Introduction)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.IsLocal) {
	orm.Where("is_local", request.IsLocal)
}
if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Path) {
	orm.Where("path", request.Path)
}
if !gconv.IsEmpty(request.DownloadCateId) {
	orm.Where("download_cate_id", request.DownloadCateId)
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


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *DownloadService) GetAll(request requests.DownloadRequest) ([]models.Download, error) {

	var list []models.Download

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Name) {
	orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Introduction) {
	orm.Where("introduction", request.Introduction)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.IsLocal) {
	orm.Where("is_local", request.IsLocal)
}
if !gconv.IsEmpty(request.AdminId) {
	orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Path) {
	orm.Where("path", request.Path)
}
if !gconv.IsEmpty(request.DownloadCateId) {
	orm.Where("download_cate_id", request.DownloadCateId)
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


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *DownloadService) Add(request requests.DownloadRequest) (bool, error) {

	var download models.Download

	download.Name = html.EscapeString(request.Name)
download.Introduction = html.EscapeString(request.Introduction)
download.Url = html.EscapeString(request.Url)
download.IsLocal = request.IsLocal
download.AdminId = request.AdminId
download.IsShow = request.IsShow
download.Sort = request.Sort
download.Path = html.EscapeString(request.Path)
download.DownloadCateId = request.DownloadCateId
download.Pic = html.EscapeString(request.Pic)
download.Platform = html.EscapeString(request.Platform)
download.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Create(&download)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *DownloadService) Save(request requests.DownloadRequest) (bool, error) {

	var download models.Download

	download.ID = request.ID
	download.Name = html.EscapeString(request.Name)
download.Introduction = html.EscapeString(request.Introduction)
download.Url = html.EscapeString(request.Url)
download.IsLocal = request.IsLocal
download.AdminId = request.AdminId
download.IsShow = request.IsShow
download.Sort = request.Sort
download.Path = html.EscapeString(request.Path)
download.DownloadCateId = request.DownloadCateId
download.Pic = html.EscapeString(request.Pic)
download.Platform = html.EscapeString(request.Platform)
download.Lang = html.EscapeString(request.Lang)


	err := facades.Orm().Query().Save(&download)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *DownloadService) Delete(id int64) (bool, error) {

	var admin models.Download
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
