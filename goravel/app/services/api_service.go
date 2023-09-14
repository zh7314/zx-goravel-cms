package services

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	"goravel/app/utils/str"
	"strings"
	"time"
)

const apiToken string = "api_token:"
const apiKey string = "api_id"

const tokenTime int = 30 * 24 * 3600

type ApiService struct {
	//Dependent services
}

func NewApiService() *ApiService {
	return &ApiService{
		//Inject model
	}
}

func (r *ApiService) GetTokenKey() string {
	return apiToken
}

func (r *ApiService) GetApiKey() string {
	return apiKey
}

func (r *ApiService) GetToken(key string, secret string) (res string, err error) {

	if len(strings.TrimSpace(key)) == 0 {
		return "", errors.New("api key不能为空")
	}
	if len(strings.TrimSpace(secret)) == 0 {
		return "", errors.New("api secret不能为空")
	}

	var apiUser models.APIUser
	err = facades.Orm.Query().Where("api_key", key).First(&apiUser)
	if err != nil {
		return "", errors.New("查询错误")
	}
	if apiUser.ID == 0 {
		return "", errors.New("key不存在")
	}
	if strings.Compare(apiUser.APISecret, secret) == -1 {
		return "", errors.New("secret不存在")
	}

	md5 := str.GetRandomString(32)

	redisKey := apiToken + md5
	err = facades.Cache.Put(redisKey, apiUser.ID, time.Duration(tokenTime)*time.Second)
	if err != nil {
		return "", errors.New("缓存失败")
	}

	return md5, nil
}
