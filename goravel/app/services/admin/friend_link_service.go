package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type FriendLinkService struct {
}

func NewFriendLinkService() *FriendLinkService {
	return &FriendLinkService{}
}

func (r *FriendLinkService) GetList(request requests.FriendLinkRequest) (map[string]interface{}, error) {

	var list []*models.FriendLink
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.IsFollow) {
		orm = orm.Where("is_follow", request.IsFollow)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
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

func (r *FriendLinkService) GetAll(request requests.FriendLinkRequest) ([]*models.FriendLink, error) {

	var list []*models.FriendLink

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.IsFollow) {
		orm = orm.Where("is_follow", request.IsFollow)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		orm = orm.Where("url", request.Url)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *FriendLinkService) GetOne(id int64) (res models.FriendLink, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var friendLink models.FriendLink
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&friendLink)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return friendLink, nil
}

func (r *FriendLinkService) Add(request requests.FriendLinkRequest) (bool, error) {

	var friendLink models.FriendLink

	if !gconv.IsEmpty(request.IsFollow) {
		friendLink.IsFollow = request.IsFollow
	}
	if !gconv.IsEmpty(request.IsShow) {
		friendLink.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		friendLink.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		friendLink.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		friendLink.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		friendLink.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Title) {
		friendLink.Title = html.EscapeString(request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		friendLink.Url = html.EscapeString(request.Url)
	}

	err := facades.Orm().Query().Create(&friendLink)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FriendLinkService) Save(request requests.FriendLinkRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var friendLink models.FriendLink
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&friendLink)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.IsFollow) {
		friendLink.IsFollow = request.IsFollow
	}
	if !gconv.IsEmpty(request.IsShow) {
		friendLink.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		friendLink.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.Pic) {
		friendLink.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Platform) {
		friendLink.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.Sort) {
		friendLink.Sort = request.Sort
	}
	if !gconv.IsEmpty(request.Title) {
		friendLink.Title = html.EscapeString(request.Title)
	}
	if !gconv.IsEmpty(request.Url) {
		friendLink.Url = html.EscapeString(request.Url)
	}

	err = facades.Orm().Query().Save(&friendLink)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FriendLinkService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var friendLink models.FriendLink
	_, err := facades.Orm().Query().Where("id", id).Delete(&friendLink)
	if err != nil {
		return false, err
	}
	return true, nil
}
