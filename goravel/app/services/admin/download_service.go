package admin

import (
	"errors"
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
	orm = orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Introduction) {
	orm = orm.Where("introduction", request.Introduction)
}
if !gconv.IsEmpty(request.Url) {
	orm = orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.IsLocal) {
	orm = orm.Where("is_local", request.IsLocal)
}
if !gconv.IsEmpty(request.AdminId) {
	orm = orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.IsShow) {
	orm = orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm = orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Path) {
	orm = orm.Where("path", request.Path)
}
if !gconv.IsEmpty(request.DownloadCateId) {
	orm = orm.Where("download_cate_id", request.DownloadCateId)
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

func (r *DownloadService) GetAll(request requests.DownloadRequest) ([]models.Download, error) {

	var list []models.Download

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Name) {
	orm = orm.Where("name", request.Name)
}
if !gconv.IsEmpty(request.Introduction) {
	orm = orm.Where("introduction", request.Introduction)
}
if !gconv.IsEmpty(request.Url) {
	orm = orm.Where("url", request.Url)
}
if !gconv.IsEmpty(request.IsLocal) {
	orm = orm.Where("is_local", request.IsLocal)
}
if !gconv.IsEmpty(request.AdminId) {
	orm = orm.Where("admin_id", request.AdminId)
}
if !gconv.IsEmpty(request.IsShow) {
	orm = orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Sort) {
	orm = orm.Where("sort", request.Sort)
}
if !gconv.IsEmpty(request.Path) {
	orm = orm.Where("path", request.Path)
}
if !gconv.IsEmpty(request.DownloadCateId) {
	orm = orm.Where("download_cate_id", request.DownloadCateId)
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


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *DownloadService) GetOne(id int64) (res models.Download, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var download models.Download
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&download)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return download, nil
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
