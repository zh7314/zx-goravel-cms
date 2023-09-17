package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
)

type AdminService struct {
}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (r *AdminService) GetList(request requests.AdminRequest) (map[string]interface{}, error) {

	var list []models.Admin
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm.Where("name like ?", "%"+request.Name+"%")
	}
	if !gconv.IsEmpty(request.ID) {
		orm.Where("id", request.ID)
	}

	orm.Order("id asc").Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *AdminService) GetAll(request requests.AdminRequest) ([]models.Admin, error) {

	var list []models.Admin

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Name) {
		orm.Where("name like ?", "%"+request.Name+"%")
	}
	if !gconv.IsEmpty(request.ID) {
		orm.Where("id", request.ID)
	}

	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *AdminService) Add(request requests.AdminRequest) (bool, error) {

	var admin models.Admin

	if !gconv.IsEmpty(request.Name) {
		admin.Name = request.Name
	}

	err := facades.Orm().Query().Create(&admin)
	if err == nil {
		return false, err
	}

	return true, nil
}

func (r *AdminService) Save(request requests.AdminRequest) (bool, error) {

	var admin models.Admin

	admin.ID = request.ID
	if !gconv.IsEmpty(request.Name) {
		admin.Name = request.Name
	}

	err := facades.Orm().Query().Save(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *AdminService) Delete(id int64) (bool, error) {

	var admin models.Admin
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
