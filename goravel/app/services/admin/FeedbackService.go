package admin

import (
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type FeedbackService struct {
}

func NewFeedbackService() *FeedbackService {
	return &FeedbackService{}
}

func (r *FeedbackService) GetList(request requests.FeedbackRequest) (map[string]interface{}, error) {

	var list []models.Feedback
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Contact) {
	orm.Where("contact", request.Contact)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.NickName) {
	orm.Where("nick_name", request.NickName)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}


	orm.Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *FeedbackService) GetAll(request requests.FeedbackRequest) ([]models.Feedback, error) {

	var list []models.Feedback

	orm := facades.Orm().Query()

    if !gconv.IsEmpty(request.Contact) {
	orm.Where("contact", request.Contact)
}
if !gconv.IsEmpty(request.Content) {
	orm.Where("content", request.Content)
}
if !gconv.IsEmpty(request.Lang) {
	orm.Where("lang", request.Lang)
}
if !gconv.IsEmpty(request.NickName) {
	orm.Where("nick_name", request.NickName)
}
if !gconv.IsEmpty(request.Platform) {
	orm.Where("platform", request.Platform)
}


	orm.Order("id desc").Get(&list)

	return list, nil
}

func (r *FeedbackService) Add(request requests.FeedbackRequest) (bool, error) {

	var feedback models.Feedback

	feedback.Contact = html.EscapeString(request.Contact)
feedback.Content = html.EscapeString(request.Content)
feedback.Lang = html.EscapeString(request.Lang)
feedback.NickName = html.EscapeString(request.NickName)
feedback.Platform = html.EscapeString(request.Platform)


	err := facades.Orm().Query().Create(&feedback)
	if err != nil {
    		return false, err
    }
	return true, nil
}

func (r *FeedbackService) Save(request requests.FeedbackRequest) (bool, error) {

	var feedback models.Feedback

	feedback.ID = request.ID
	feedback.Contact = html.EscapeString(request.Contact)
feedback.Content = html.EscapeString(request.Content)
feedback.Lang = html.EscapeString(request.Lang)
feedback.NickName = html.EscapeString(request.NickName)
feedback.Platform = html.EscapeString(request.Platform)


	err := facades.Orm().Query().Save(&feedback)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FeedbackService) Delete(id int64) (bool, error) {

	var admin models.Feedback
	_, err := facades.Orm().Query().Delete(&admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
