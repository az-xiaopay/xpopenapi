package xpopenapi

import (
	"xpopenapi/common/config"
)

type openApi struct {
}

func New(env Env, appId, appSecret, appKey string) openApi {
	if env != Env_Dev && env != Env_Prod {
		panic("err env")
	}
	config.Base.Env = string(env)
	if env == Env_Dev {
		config.Base.XPServerUrl = "https://fpdev.xiaopay.net"
	} else {
		config.Base.XPServerUrl = "https://pay.xiaopay.net"
	}
	config.OpenApp.AppId = appId
	config.OpenApp.AppSecret = appSecret
	config.OpenApp.AppKey = appKey

	return openApi{}
}
