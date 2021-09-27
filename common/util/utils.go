package util

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	mathrand "math/rand"
	"regexp"
	"strings"
	"time"
)

const Random_phone_prefix = "101"

// CheckPhone 验证手机号码
func CheckPhone(mobileNum string) bool {
	sregular := "^1[3456789]\\d{9}$"
	reg := regexp.MustCompile(sregular)
	return reg.MatchString(mobileNum)
}

//MD5 md5加密
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

const tokenGenKey = "hfdsijahflksdhf"

func CreateToken(s string) string {
	s = strings.ReplaceAll(s, "|", "")
	return s + "|" + MD5(s+tokenGenKey)
}
func ParseToken(s string) string {
	a := strings.Split(s, "|")
	if len(a) != 2 {
		return ""
	}
	if MD5(a[0]+tokenGenKey) != a[1] {
		return ""
	}
	return a[0]
}

// 随机生成号码
func GenerateRandomPhone(prefix string) string {
	return prefix + fmt.Sprintf("%08v", mathrand.New(mathrand.NewSource(time.Now().UnixNano())).Int63n(100000000))
}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//GetGUID 产生GUID
func GetGUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}
