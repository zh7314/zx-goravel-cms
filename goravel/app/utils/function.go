package utils

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/local"
	"goravel/app/utils/str"
	"hash/crc32"
	"strings"
	"time"
)

//func TreeMenu[T any](list []T, parentId int64) []T {
//
//	data := make([]T, 0)
//	for _, v := range list {
//		if v.ParentId == parentId {
//			//v.Label = v.Name
//			v.Children = TreeMenu(list, v.ID)
//			data = append(data, v)
//		}
//	}
//	return data
//}

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

func CRC32(input string) uint32 {
	bytes := []byte(input)
	return crc32.ChecksumIEEE(bytes)
}

func GetUniqid() string {
	id := uuid.New()
	return str.Md5(id.String())
}

func GetIP(ctx http.Context) string {

	ip := strings.TrimSpace(ctx.Request().Header("X-Forwarded-For", "127.0.0.1"))
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(ctx.Request().Header("X-Real-Ip", "127.0.0.1"))
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(ctx.Request().Ip())
	if ip != "" {
		return ip
	}
	return "127.0.0.1"
}

func TimeToLocalTime(t time.Time) (local.LocalTime, error) {
	return local.LocalTime(t), nil
}

func LocalTimeToTime(t local.LocalTime) (r time.Time, err error) {

	t1, err := time.Parse("2006-01-02 15:04:05", t.String())
	if err != nil {
		return r, err
	}
	return t1, nil
}

func StrToLocalTime(t string) (r local.LocalTime, err error) {

	t1, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return r, err
	}
	return local.LocalTime(t1), nil
}
