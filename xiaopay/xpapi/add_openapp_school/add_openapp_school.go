package add_openapp_school

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
	Schoolid int64 `json:"schoolid"` //学校ID
}

//添加开放应用学校权限
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_AddOpenAppSchool, access_token), req)
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
