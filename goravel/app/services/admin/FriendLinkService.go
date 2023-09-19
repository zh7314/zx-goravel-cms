package admin

import (
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

	var list []models.FriendLink
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.IsFollow) {
	orm.Where("is_follow", request.IsFollow)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
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
if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *FriendLinkService) GetAll(request requests.FriendLinkRequest) ([]models.FriendLink, error) {

	var list []models.FriendLink

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.IsFollow) {
	orm.Where("is_follow", request.IsFollow)
}
if !gconv.IsEmpty(request.IsShow) {
	orm.Where("is_show", request.IsShow)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
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
if !gconv.IsEmpty(request.Title) {
	orm.Where("title", request.Title)
}
if !gconv.IsEmpty(request.Url) {
	orm.Where("url", request.Url)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *FriendLinkService) Add(request requests.FriendLinkRequest) (bool, error) {

	var friendLink models.FriendLink

	friendLink.IsFollow = request.IsFollow
friendLink.IsShow = request.IsShow
friendLink.Lang = html.EscapeString(request.Lang)
friendLink.Pic = html.EscapeString(request.Pic)
friendLink.Platform = html.EscapeString(request.Platform)
friendLink.Sort = request.Sort
friendLink.Title = html.EscapeString(request.Title)
friendLink.Url = html.EscapeString(request.Url)


	err := facades.Orm().Query().Create(&friendLink)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *FriendLinkService) Save(request requests.FriendLinkRequest) (bool, error) {

	var friendLink models.FriendLink

	friendLink.ID = request.ID
	friendLink.IsFollow = request.IsFollow
friendLink.IsShow = request.IsShow
friendLink.Lang = html.EscapeString(request.Lang)
friendLink.Pic = html.EscapeString(request.Pic)
friendLink.Platform = html.EscapeString(request.Platform)
friendLink.Sort = request.Sort
friendLink.Title = html.EscapeString(request.Title)
friendLink.Url = html.EscapeString(request.Url)


	err := facades.Orm().Query().Save(&friendLink)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FriendLinkService) Delete(id int64) (bool, error) {

	var admin models.FriendLink
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
