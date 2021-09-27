package toggle_facepay

import (
	"encoding/json"
	"fmt"
	"xpopenapi/common/config"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
	"xpopenapi/xiaopay/signature"
)

type Resp struct {
	xiaopay.RespBase
	Data string `json:"Data"`
}

type PubParame struct {
	Appid    string `json:"appid"`
	Noncestr string `json:"noncestr"`
	Signtype string `json:"signtype"`
	Sign     string `json:"sign"`
}

type Req struct {
	PubParame
	Schoolid int64 `json:"schoolid"` //学校ID
	Userid   int64 `json:"userid"`   //用户id
	Status   int64 `json:"status"`   //开启状态 1开启  0关闭
}

//设置刷脸开关状态
//access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	req.Appid = config.OpenApp.AppId
	req.Noncestr = "hello"
	req.Signtype = signature.MD5
	req.Sign = signature.Sign(signature.StructToParams(req, "Req"), req.Signtype, config.OpenApp.AppKey)
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_DelSchoolClass, access_token), req)
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
