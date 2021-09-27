package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func ToInt64(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func ToInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}

func ToStr(v interface{}) string {
	switch v.(type) {
	case int:
		return strconv.Itoa(v.(int))
	case int64:
		return strconv.FormatInt(v.(int64), 10)
	case string:
		return v.(string)
	}
	return ""
}

func ToJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func ToIds(ids interface{}) string {
	ret := ToJson(ids)
	ret = strings.Replace(ret, "[", "", 1)
	ret = strings.Replace(ret, "]", "", 1)
	return ret
}

func CnNum(num int) string {
	const name = "一二三四五六七八九"
	if num < 1 {
		num = 1
	}
	if num > 9 {
		num = 9
	}
	return name[(num-1)*3 : num*3]
}

func CnNumExceed(num int) string {
	gradeName := "十"
	switch num {
	case 0:
		gradeName = "十一"
	case 1:
		gradeName = "十二"
	case 2:
		gradeName = "十三"
	case 3:
		gradeName = "十"
	}
	return gradeName
}

func CnGetKindergarten(num int) string {

	gradeName := "幼儿园小小班"
	switch num {
	case 0:
		gradeName = "非标准年级"
	case 1:
		gradeName = "幼儿园小小班"
	case 2:
		gradeName = "幼儿园小班"
	case 3:
		gradeName = "幼儿园中班"
	case 4:
		gradeName = "幼儿园大班"
	case 5:
		gradeName = "幼儿园学前班"
	default:
		gradeName = "幼儿园小小班"
	}
	return gradeName
}

func Serialize(slice []int) string {

	str := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ";", -1)

	return str
}

func Unserialize(str string) []int {

	retslice := make([]int, 0)
	slice := strings.Split(str, ";")
	for _, s := range slice {
		retslice = append(retslice, ToInt(s))
	}
	return retslice
}
