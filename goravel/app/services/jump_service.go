package services

import (
	"errors"
	"github.com/goravel/framework/facades"
	"goravel/app/models"
	"goravel/app/utils"
	"goravel/app/utils/local"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const jumpKey string = "JumpLink:"

type JumpService struct {
	//Dependent services
}

func NewJumpService() *JumpService {
	return &JumpService{
		//Inject model
	}
}

func (r *JumpService) GetJumpKey() string {
	return jumpKey
}

func (r *JumpService) GetData() []models.JumpLink {
	var jump []models.JumpLink
	facades.Orm.Query().Get(&jump)

	return jump
}

func (r *JumpService) DoJump(id string) (jumpUrl string, ok error) {

	if len(strings.TrimSpace(id)) == 0 {
		return "", errors.New("查询条件不能为空")
	}

	key := jumpKey + id
	isHas := facades.Cache.Has(key)
	if !isHas {
		return "", errors.New("查询key不存在")
	}

	res := facades.Cache.GetString(key)
	if len(res) == 0 {
		return "", errors.New("查询值为空")
	}

	jump, err := url.QueryUnescape(res)
	if err != nil {
		return "", err
	}

	return jump, nil
}

func (r *JumpService) AddLink(url_str string, end_time string) (res bool, ok error) {

	if len(strings.TrimSpace(url_str)) == 0 {
		return false, errors.New("url不能为空")
	}
	if len(strings.TrimSpace(end_time)) == 0 {
		return false, errors.New("end_time不能为空")
	}

	now := time.Now()
	start_time := now.Unix()

	end, err := time.Parse("2006-01-02 15:04:05", end_time)
	if err != nil {
		return false, errors.New("解析结束时间错误")
	}

	end_time_unix := end.Unix()
	if start_time > end_time_unix {
		return false, errors.New("结束时间不能大于当前时间")
	}

	count := end_time_unix - start_time

	var jump models.JumpLink
	jump.JumpURL = url.QueryEscape(url_str)

	jump.EndTime = local.LocalTime(end)

	err1 := facades.Orm.Query().Save(&jump)
	if err1 != nil {
		return false, errors.New("保存失败")
	}

	str_id := strconv.FormatInt(int64(jump.ID), 10)
	m_id := utils.CRC32(str_id)
	jump.HashKey = strconv.FormatInt(int64(m_id), 10)

	err2 := facades.Orm.Query().Save(&jump)
	if err2 != nil {
		return false, err2
	}

	key := jumpKey + jump.HashKey
	err3 := facades.Cache.Put(key, jump.JumpURL, time.Duration(count)*time.Second)
	if err3 != nil {
		return false, err3
	}

	return true, nil
}

func (r *JumpService) ScanLink() error {

	now := time.Now()
	end := now.Format("2006-01-02 15:04:05")

	var count int64
	facades.Orm.Query().Model(&models.JumpLink{}).Where("end_time > ?", end).Count(&count)

	if count > 0 {
		var jumps []models.JumpLink
		facades.Orm.Query().Where("end_time > ?", end).Get(&jumps)

		for _, jump := range jumps {

			key := jumpKey + jump.HashKey
			//fmt.Println(key)

			isHas := facades.Cache.Has(key)
			// 没有数据就重新写入到redis
			if !isHas {
				now := time.Now()
				start_time := now.Unix()

				end, err := time.Parse("2006-01-02 15:04:05", jump.EndTime.String())
				if err != nil {
					return err
				}

				end_time_unix := end.Unix()
				if start_time > end_time_unix {
					return errors.New("结束时间不能大于当前时间")
				}

				count := end_time_unix - start_time
				err = facades.Cache.Put(key, jump.JumpURL, time.Duration(count)*time.Second)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}
