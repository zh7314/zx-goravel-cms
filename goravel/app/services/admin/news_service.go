package admin

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	requests "goravel/app/requests/admin"
	"goravel/app/utils/gconv"
	"html"
)

type NewsService struct {
}

func NewNewsService() *NewsService {
	return &NewsService{}
}

func (r *NewsService) GetList(request requests.NewsRequest) (map[string]interface{}, error) {

	var list []*models.News
	var count int64

	orm := facades.Orm().Query().With("NewsCate")

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !gconv.IsEmpty(request.Count) {
		orm = orm.Where("count", request.Count)
	}
	if !request.EndTime.IsZero() {
		orm = orm.Where("end_time", request.EndTime)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.NewsCateId) {
		orm = orm.Where("news_cate_id", request.NewsCateId)
	}
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		orm = orm.Where("pics", request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.SeoDescription) {
		orm = orm.Where("seo_description", request.SeoDescription)
	}
	if !gconv.IsEmpty(request.SeoKeyword) {
		orm = orm.Where("seo_keyword", request.SeoKeyword)
	}
	if !gconv.IsEmpty(request.SeoTitle) {
		orm = orm.Where("seo_title", request.SeoTitle)
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		orm = orm.Where("short_title", request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !request.StartTime.IsZero() {
		orm = orm.Where("start_time", request.StartTime)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}

	if request.Page > 0 && request.PageSize > 0 {
		orm.Order("sort asc").Order("id desc").Paginate(request.Page, request.PageSize, &list, &count)
	} else {
		orm.Order("sort asc").Order("id desc").Get(&list)
		count = int64(len(list))
	}

	if len(list) > 0 {
		for _, v := range list {
			v.Content = html.UnescapeString(v.Content)
		}
	}

	res := make(map[string]interface{})
	res["list"] = list
	res["count"] = count

	return res, nil
}

func (r *NewsService) GetAll(request requests.NewsRequest) ([]*models.News, error) {

	var list []*models.News

	orm := facades.Orm().Query()

	if !gconv.IsEmpty(request.AdminId) {
		orm = orm.Where("admin_id", request.AdminId)
	}
	if !gconv.IsEmpty(request.Content) {
		orm = orm.Where("content", request.Content)
	}
	if !gconv.IsEmpty(request.Count) {
		orm = orm.Where("count", request.Count)
	}
	if !request.EndTime.IsZero() {
		orm = orm.Where("end_time", request.EndTime)
	}
	if !gconv.IsEmpty(request.IsShow) {
		orm = orm.Where("is_show", request.IsShow)
	}
	if !gconv.IsEmpty(request.Lang) {
		orm = orm.Where("lang", request.Lang)
	}
	if !gconv.IsEmpty(request.NewsCateId) {
		orm = orm.Where("news_cate_id", request.NewsCateId)
	}
	if !gconv.IsEmpty(request.Pic) {
		orm = orm.Where("pic", request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		orm = orm.Where("pics", request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		orm = orm.Where("platform", request.Platform)
	}
	if !gconv.IsEmpty(request.SeoDescription) {
		orm = orm.Where("seo_description", request.SeoDescription)
	}
	if !gconv.IsEmpty(request.SeoKeyword) {
		orm = orm.Where("seo_keyword", request.SeoKeyword)
	}
	if !gconv.IsEmpty(request.SeoTitle) {
		orm = orm.Where("seo_title", request.SeoTitle)
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		orm = orm.Where("short_title", request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		orm = orm.Where("sort", request.Sort)
	}
	if !request.StartTime.IsZero() {
		orm = orm.Where("start_time", request.StartTime)
	}
	if !gconv.IsEmpty(request.Title) {
		orm = orm.Where("title", request.Title)
	}

	orm.Order("sort asc").Order("id desc").Get(&list)

	return list, nil
}

func (r *NewsService) GetOne(id int64) (res models.News, err error) {

	if gconv.IsEmpty(id) {
		return res, errors.New("id不能为空")
	}

	var news models.News
	err = facades.Orm().Query().Where("id", id).FirstOrFail(&news)
	if err != nil {
		return res, errors.New("数据不存在")
	}

	return news, nil
}

func (r *NewsService) Add(request requests.NewsRequest) (bool, error) {

	var news models.News

	if !gconv.IsEmpty(request.AdminId) {
		news.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.Content) {
		news.Content = html.EscapeString(request.Content)
	}
	if !gconv.IsEmpty(request.Count) {
		news.Count = request.Count
	}
	if !request.EndTime.IsZero() {
		news.EndTime = &request.EndTime
	}
	if !gconv.IsEmpty(request.IsShow) {
		news.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		news.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.NewsCateId) {
		news.NewsCateId = request.NewsCateId
	}
	if !gconv.IsEmpty(request.Pic) {
		news.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		news.Pics = html.EscapeString(request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		news.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.SeoDescription) {
		news.SeoDescription = html.EscapeString(request.SeoDescription)
	}
	if !gconv.IsEmpty(request.SeoKeyword) {
		news.SeoKeyword = html.EscapeString(request.SeoKeyword)
	}
	if !gconv.IsEmpty(request.SeoTitle) {
		news.SeoTitle = html.EscapeString(request.SeoTitle)
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		news.ShortTitle = html.EscapeString(request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		news.Sort = request.Sort
	}
	if !request.StartTime.IsZero() {
		news.StartTime = &request.StartTime
	}
	if !gconv.IsEmpty(request.Title) {
		news.Title = html.EscapeString(request.Title)
	}

	err := facades.Orm().Query().Create(&news)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *NewsService) Save(request requests.NewsRequest) (bool, error) {

	if gconv.IsEmpty(request.ID) {
		return false, errors.New("请求不能为空")
	}

	var news models.News
	err := facades.Orm().Query().Where("id", request.ID).FirstOrFail(&news)
	if err != nil {
		return false, errors.New("数据不存在")
	}

	if !gconv.IsEmpty(request.AdminId) {
		news.AdminId = request.AdminId
	}
	if !gconv.IsEmpty(request.Content) {
		news.Content = html.EscapeString(request.Content)
	}
	if !gconv.IsEmpty(request.Count) {
		news.Count = request.Count
	}
	if !request.EndTime.IsZero() {
		news.EndTime = &request.EndTime
	}
	if !gconv.IsEmpty(request.IsShow) {
		news.IsShow = request.IsShow
	}
	if !gconv.IsEmpty(request.Lang) {
		news.Lang = html.EscapeString(request.Lang)
	}
	if !gconv.IsEmpty(request.NewsCateId) {
		news.NewsCateId = request.NewsCateId
	}
	if !gconv.IsEmpty(request.Pic) {
		news.Pic = html.EscapeString(request.Pic)
	}
	if !gconv.IsEmpty(request.Pics) {
		news.Pics = html.EscapeString(request.Pics)
	}
	if !gconv.IsEmpty(request.Platform) {
		news.Platform = html.EscapeString(request.Platform)
	}
	if !gconv.IsEmpty(request.SeoDescription) {
		news.SeoDescription = html.EscapeString(request.SeoDescription)
	}
	if !gconv.IsEmpty(request.SeoKeyword) {
		news.SeoKeyword = html.EscapeString(request.SeoKeyword)
	}
	if !gconv.IsEmpty(request.SeoTitle) {
		news.SeoTitle = html.EscapeString(request.SeoTitle)
	}
	if !gconv.IsEmpty(request.ShortTitle) {
		news.ShortTitle = html.EscapeString(request.ShortTitle)
	}
	if !gconv.IsEmpty(request.Sort) {
		news.Sort = request.Sort
	}
	if !request.StartTime.IsZero() {
		news.StartTime = &request.StartTime
	}
	if !gconv.IsEmpty(request.Title) {
		news.Title = html.EscapeString(request.Title)
	}

	err = facades.Orm().Query().Save(&news)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *NewsService) Delete(id int64) (bool, error) {

	if gconv.IsEmpty(id) {
		return false, errors.New("id不能为空")
	}

	var news models.News
	_, err := facades.Orm().Query().Where("id", id).Delete(&news)
	if err != nil {
		return false, err
	}
	return true, nil
}
