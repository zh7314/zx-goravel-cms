package str

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/axgle/mahonia"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unsafe"
)

func Md5(buf string) string {

	data := []byte(buf)
	hash := md5.New()
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// 用b的所有字段覆盖a的
// 如果fields不为空, 表示用b的特定字段覆盖a的
// a应该为结构体指针
func CopyFields(a interface{}, b interface{}, fields ...string) (err error) {
	at := reflect.TypeOf(a)
	av := reflect.ValueOf(a)
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)
	// 简单判断下
	if at.Kind() != reflect.Ptr {
		err = fmt.Errorf("a must be a struct pointer")
		return
	}
	av = reflect.ValueOf(av.Interface())
	// 要复制哪些字段
	_fields := make([]string, 0)
	if len(fields) > 0 {
		_fields = fields
	} else {
		for i := 0; i < bv.NumField(); i++ {
			_fields = append(_fields, bt.Field(i).Name)
		}
	}
	if len(_fields) == 0 {
		fmt.Println("no fields to copy")
		return
	}
	// 复制
	for i := 0; i < len(_fields); i++ {
		name := _fields[i]
		f := av.Elem().FieldByName(name)
		bValue := bv.FieldByName(name)
		// a中有同名的字段并且类型一致才复制
		if f.IsValid() && f.Kind() == bValue.Kind() {
			f.Set(bValue)
		} else {
			//fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}

// 结构体转map gorm Updates不会更新结构体里的0及空，需转换成map
// 此方法不适用此情景，更新为Struct2MapByTag
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// 数组转字符串 逗号分割
func Array2Str(s interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), " ", ",", -1)
}

func GetBase64ByFile(path string) (string, error) {
	ff, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	return sourcestring, nil
}

func Struct2MapByTag(obj interface{}, tagName string) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(tagName) == "" {
			continue
		}
		data[t.Field(i).Tag.Get(tagName)] = v.Field(i).Interface()
	}
	return data
}

// 判断字符串是否在数组里
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func SetPassword(len int, pwdO string) (pwd string, salt string) {
	salt = GetRandomString(len)
	defaultPwd := pwdO
	pwd = Md5(defaultPwd + salt)
	return pwd, salt
}

// 生成随机字符串
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func IsEmail(b []byte) bool {
	var emailPattern = regexp.MustCompile("[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[a-zA-Z0-9](?:[\\w-]*[\\w])?")
	return emailPattern.Match(b)
}

func JobKey(taskId, serverId int) int {
	return taskId*100000 + serverId
}

func GbkAsUtf8(str string) string {
	srcDecoder := mahonia.NewDecoder("gbk")
	desDecoder := mahonia.NewDecoder("utf-8")
	resStr := srcDecoder.ConvertString(str)
	_, resBytes, _ := desDecoder.Translate([]byte(resStr), true)
	return string(resBytes)
}

// FromBytes converts the specified byte array to a string.
func FromBytes(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// Bytes converts the specified str to a byte array.
func ToBytes(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// LCS gets the longest common substring of s1 and s2.
//
// Refers to http://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Longest_common_substring.
func LCS(s1 string, s2 string) string {
	var m = make([][]int, 1+len(s1))

	for i := 0; i < len(m); i++ {
		m[i] = make([]int, 1+len(s2))
	}

	longest := 0
	xLongest := 0

	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				m[x][y] = m[x-1][y-1] + 1
				if m[x][y] > longest {
					longest = m[x][y]
					xLongest = x
				}
			} else {
				m[x][y] = 0
			}
		}
	}
	return s1[xLongest-longest : xLongest]
}
