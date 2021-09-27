package add_openapp_interface

import (
	"encoding/json"
	"errors"
	"fmt"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
)

type Resp struct {
	xiaopay.RespBase
	Data string `json:"Data"`
}

type Req struct {
	Interfaceid int64 `json:"interfaceid"` //接口ID
}

//添加开放应用学校权限
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_AddOpenAppInterface, access_token), req)
	if err != nil {
		return nil, err
	}
	var resp Resp
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Success {
		return &resp, errors.New(resp.ErrorMessage)
	}

	return &resp, nil
}
