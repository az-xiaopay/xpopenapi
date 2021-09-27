package query_user_school

import (
	"encoding/json"
	"errors"
	"fmt"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
)

type Resp struct {
	xiaopay.RespBase
	Data Data `json:"Data"`
}

type Data struct {
	Phone    string `json:"phone"`
	Schoolid int64  `json:"schoolid"`
	Userid   int64  `json:"userid"`
	Username string `json:"username"`
}

type Req struct {
	Schoolid int64  `json:"schoolid"`
	Phone    string `json:"phone"`
}

//查询学校班级信息（单个）
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_QueryUserSchool, access_token), req)
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
