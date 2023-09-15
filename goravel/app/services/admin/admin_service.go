package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
)

type AdminService struct {
	//Dependent services
}

func NewAdminService() *AdminService {
	return &AdminService{
		//Inject model
	}
}
func (r *AdminService) GetList() ([]models.Banner, error) {

	//var admin models.Admin
	//facades.Orm().Query().With("Author").With("Publisher").Find(&admin)

	var banner []models.Banner

	facades.Orm().Query().With("Admin").Get(&banner)

	return banner, nil
}

func (r *AdminService) GetAll() (int, error) {

	return 12354, nil
}
func (r *AdminService) Add() (int, error) {

	return 12354, nil
}
func (r *AdminService) Save() (int, error) {

	return 12354, nil
}
func (r *AdminService) Delete() (int, error) {

	return 12354, nil
}
