package xptool

import (
	"sync"
	"time"
	"xpopenapi/common/config"
	"xpopenapi/xiaopay/xpapi/application_token"
)

var access_token accessToken
var lock sync.RWMutex

func GetAccessToken() (string, error) {
	lock.Lock()
	defer lock.Unlock()
	ct := access_token
	if ct.expire_timestamp < time.Now().Unix() {
		do, err := application_token.Do(config.OpenApp.AppId, config.OpenApp.AppSecret)
		if err != nil {
			return "", err
		}
		access_token = accessToken{
			token:            do.Data.AccessToken,
			expire_timestamp: time.Now().Unix() - 100 + int64(do.Data.ExpiresIn),
		}
		return access_token.token, nil
	}
	return ct.token, nil
}

type accessToken struct {
	token            string
	expire_timestamp int64
}
