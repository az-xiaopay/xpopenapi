package signature

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

const (
	HMACSHA256 = "HMAC-SHA256"
	MD5        = "MD5"
)

// 开放平台专用签名方法,请知悉，无移动使用
func Sign(params Params, signType, apikey string) string {
	// 创建切片
	var keys = make([]string, 0, len(params))
	// 遍历签名参数
	for k := range params {
		if k != "Sign" { // 排除sign字段
			keys = append(keys, k)
		}
	}
	// 由于切片的元素顺序是不固定，所以这里强制给切片元素加个顺序
	sort.Strings(keys)

	//创建字符缓冲
	var buf bytes.Buffer
	for _, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(strings.ToLower(k))
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			buf.WriteString(`&`)
		}
	}
	fmt.Println(buf.String())

	// 加入apiKey作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(apikey)

	fmt.Println(buf.String())

	var (
		dataMd5    [16]byte
		dataSha256 []byte
		str        string
	)

	switch signType {
	case MD5:
		dataMd5 = md5.Sum(buf.Bytes())
		str = hex.EncodeToString(dataMd5[:]) //需转换成切片
	case HMACSHA256:
		h := hmac.New(sha256.New, []byte(apikey))
		h.Write(buf.Bytes())
		dataSha256 = h.Sum(nil)
		str = hex.EncodeToString(dataSha256[:])
	}
	return strings.ToUpper(str)
}

func MapToXml(params map[string]interface{}) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		if v != nil {
			buf.WriteString(`<`)
			buf.WriteString(k)
			buf.WriteString(`><![CDATA[`)

			switch v.(type) {
			case string:
				buf.WriteString(v.(string))
			default:
				buf.WriteString("")
			}

			buf.WriteString(v.(string))
			buf.WriteString(`]]></`)
			buf.WriteString(k)
			buf.WriteString(`>`)
		}

	}
	buf.WriteString(`</xml>`)
	return buf.String()
}
