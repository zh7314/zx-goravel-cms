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

	var list []*models.File
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.FileName) {
		orm = orm.Where("file_name", request.FileName)
	}
	if !gconv.IsEmpty(request.FilePath) {
		orm = orm.Where("file_path", request.FilePath)
	}
	if !gconv.IsEmpty(request.FileSize) {
		orm = orm.Where("file_size", request.FileSize)
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

func (r *FileService) GetAll(request requests.FileRequest) ([]*models.File, error) {

	var list []*models.File

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.FileName) {
		orm = orm.Where("file_name", request.FileName)
	}
	if !gconv.IsEmpty(request.FilePath) {
		orm = orm.Where("file_path", request.FilePath)
	}
	if !gconv.IsEmpty(request.FileSize) {
		orm = orm.Where("file_size", request.FileSize)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

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

	if !gconv.IsEmpty(request.AdminId) {
		file.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.FileName) {
		file.FileName = html.EscapeString(request.FileName)
	}
	if !gconv.IsEmpty(request.FilePath) {
		file.FilePath = html.EscapeString(request.FilePath)
	}
	if !gconv.IsEmpty(request.FileSize) {
		file.FileSize = request.FileSize
	}

	err := facades.Orm().Query().Create(&file)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FileService) Save(request requests.FileRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var file models.File
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&file)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.AdminId) {
		file.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.FileName) {
		file.FileName = html.EscapeString(request.FileName)
	}
	if !gconv.IsEmpty(request.FilePath) {
		file.FilePath = html.EscapeString(request.FilePath)
	}
	if !gconv.IsEmpty(request.FileSize) {
		file.FileSize = request.FileSize
	}

	err = facades.Orm().Query().Save(&file)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FileService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var file models.File
	_, err := facades.Orm().Query().Where("id", id).Delete(&file)
	if err != nil {
		return false, err
	}
	return true, nil
}
