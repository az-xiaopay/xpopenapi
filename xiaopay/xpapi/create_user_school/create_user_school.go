package create_user_school

import (
	"encoding/json"
	"fmt"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
)

type Resp struct {
	xiaopay.RespBase
	SchoolUserId int64 `json:"Data"`
}

type Req struct {
	Schoolid int64  `json:"schoolid"` //学校ID
	Phone    string `json:"phone"`    //手机号
	Name     string `json:"name"`     //管理员姓名
}

//创建学校管理员
//get_access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_CreateUserSchool, access_token), req)
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
