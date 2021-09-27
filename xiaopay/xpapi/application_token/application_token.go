package application_token

import (
	"encoding/json"
	"fmt"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
)

type Resp struct {
	xiaopay.RespBase
	Data Data `json:"Data"`
}
type Data struct {
	AccessToken string `json:"access_token"` //接口访问令牌
	ExpiresIn   int    `json:"expires_in"`   //令牌过期时间(秒)
}

//获取Access_token
//该接口用于校智付开放平台应用根据appid和appsecret获取access_token，appid和appsecret随对接开放平台时应用创建生成
//appid 开放应用唯一标识
//appsecret 开放应用秘钥
func Do(appid string, appsecret string) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_application_token), struct {
		Appid     string `json:"appid"`
		Appsecret string `json:"appsecret"`
	}{appid, appsecret})
	if err != nil {
		return nil, err
	}
	var resp Resp
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
