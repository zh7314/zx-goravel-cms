package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (r *FileService) GetList(request requests.FileRequest) (map[string]interface{}, error) {

	var list []models.File
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.FileName) {
	orm = orm.Where("file_name", request.FileName)
}
if !gconv.IsEmpty(request.FilePath) {
	orm = orm.Where("file_path", request.FilePath)
}
if !gconv.IsEmpty(request.FileSize) {
	orm = orm.Where("file_size", request.FileSize)
}
if !gconv.IsEmpty(request.AdminId) {
	orm = orm.Where("admin_id", request.AdminId)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *FileService) GetAll(request requests.FileRequest) ([]models.File, error) {

	var list []models.File

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.FileName) {
	orm = orm.Where("file_name", request.FileName)
}
if !gconv.IsEmpty(request.FilePath) {
	orm = orm.Where("file_path", request.FilePath)
}
if !gconv.IsEmpty(request.FileSize) {
	orm = orm.Where("file_size", request.FileSize)
}
if !gconv.IsEmpty(request.AdminId) {
	orm = orm.Where("admin_id", request.AdminId)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *FileService) GetOne(id int64) (res models.File, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var file models.File
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&file)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return file, nil
}

func (r *FileService) Add(request requests.FileRequest) (bool, error) {

	var file models.File

	file.FileName = html.EscapeString(request.FileName)
file.FilePath = html.EscapeString(request.FilePath)
file.FileSize = request.FileSize
file.AdminId = request.AdminId


	err := facades.Orm().Query().Create(&file)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *FileService) Save(request requests.FileRequest) (bool, error) {

	var file models.File

	file.ID = request.ID
	file.FileName = html.EscapeString(request.FileName)
file.FilePath = html.EscapeString(request.FilePath)
file.FileSize = request.FileSize
file.AdminId = request.AdminId


	err := facades.Orm().Query().Save(&file)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FileService) Delete(id int64) (bool, error) {

	var admin models.File
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
