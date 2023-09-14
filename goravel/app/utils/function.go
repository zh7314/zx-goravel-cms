package utils

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/utils/local"
	"hash/crc32"
	"strings"
	"time"
)

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

func TimeToLocalTime(t time.Time) (res local.LocalTime, err error) {
	return local.LocalTime(t), nil
}

func LocalTimeToTime(t local.LocalTime) (res time.Time, err error) {

	time, err := time.Parse("2006-01-02 15:04:05", t.String())
	if err != nil {
		return res, err
	}
	return time, nil
}
