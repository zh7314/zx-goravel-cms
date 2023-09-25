package admin

import "fmt"

type CommonService struct {
}

func NewCommonService() *CommonService {
	return &CommonService{}
}

func (r *CommonService) Check(adminId int64, url string) error {
	fmt.Println(adminId)
	fmt.Println(url)

	return nil
}
