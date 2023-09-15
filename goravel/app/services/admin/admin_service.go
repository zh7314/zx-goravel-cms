package admin

import (
	"github.com/goravel/framework/facades"
	models "goravel/app/http/requests/admin"
)

type AdminService struct {
	//Dependent services
}

func NewAdminService() *AdminService {
	return &AdminService{
		//Inject model
	}
}
func (r *AdminService) GetList() (int, error) {

	var admin models.Admin
	facades.Orm().Query().With("Author").With("Publisher").Find(&admin)

	return 12354, nil
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
