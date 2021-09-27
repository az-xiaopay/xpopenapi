package config

type base struct {
	XPServerUrl string
	Env         string
}

type openapp struct {
	AppId     string
	AppSecret string
	AppKey    string
}

var Base base
var OpenApp openapp
