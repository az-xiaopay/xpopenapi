package get_weapp_free_login_jump

import (
	"encoding/json"
	"fmt"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
)

type Resp struct {
	xiaopay.RespBase
	Data interface{} `json:"Data"`
}

type Req struct {
	Schoolid int64 `json:"schoolid"` //学校ID
	Userid   int64 `json:"userid"`   //平台用户id
}

//获取免登录跳转地址
//access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_WeappFreeLoginJump, access_token), req)
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
