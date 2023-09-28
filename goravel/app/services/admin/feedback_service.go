package admin

import (
	"errors"
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

	var list []*models.Feedback
	var count int64

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Contact) {
		orm = orm.Where("contact", request.Contact)
	}
	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.NickName) {
		orm = orm.Where("nick_name", request.NickName)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
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

func (r *FeedbackService) GetAll(request requests.FeedbackRequest) ([]*models.Feedback, error) {

	var list []*models.Feedback

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.Contact) {
		orm = orm.Where("contact", request.Contact)
	}
	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.NickName) {
		orm = orm.Where("nick_name", request.NickName)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *FeedbackService) GetOne(id int64) (res models.Feedback, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var feedback models.Feedback
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&feedback)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return feedback, nil
}

func (r *FeedbackService) Add(request requests.FeedbackRequest) (bool, error) {

	var feedback models.Feedback

	if !gconv.IsEmpty(request.Contact) {
		feedback.Contact = html.EscapeString(request.Contact)
	}
	if !gconv.IsEmpty(request.Content) {
		feedback.Content = html.EscapeString(request.Content)
	}
	if !gconv.IsEmpty(request.Lang) {
		feedback.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.NickName) {
		feedback.NickName = html.EscapeString(request.NickName)
	}
	if !gconv.IsEmpty(request.Platform) {
		feedback.Platform = html.EscapeString(request.Platform)
	}

	err := facades.Orm().Query().Create(&feedback)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FeedbackService) Save(request requests.FeedbackRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var feedback models.Feedback
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&feedback)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.Contact) {
		feedback.Contact = html.EscapeString(request.Contact)
	}
	if !gconv.IsEmpty(request.Content) {
		feedback.Content = html.EscapeString(request.Content)
	}
	if !gconv.IsEmpty(request.Lang) {
		feedback.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.NickName) {
		feedback.NickName = html.EscapeString(request.NickName)
	}
	if !gconv.IsEmpty(request.Platform) {
		feedback.Platform = html.EscapeString(request.Platform)
	}

	err = facades.Orm().Query().Save(&feedback)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *FeedbackService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var feedback models.Feedback
	_, err := facades.Orm().Query().Where("id", id).Delete(&feedback)
	if err != nil {
		return false, err
	}
	return true, nil
}
