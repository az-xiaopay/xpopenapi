package query_school_user

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
	Cardnumber           string `json:"cardnumber"`
	Classid              int64  `json:"classid"`
	Classname            string `json:"classname"`
	Gender               int    `json:"gender"`
	Gradeid              int64  `json:"gradeid"`
	Gradename            string `json:"gradename"`
	Identificationnumber string `json:"identificationnumber"`
	Identificationtype   string `json:"identificationtype"`
	Phonenumber          string `json:"phonenumber"`
	Userid               int64  `json:"userid"`
	Username             string `json:"username"`
	Usertype             string `json:"usertype"`
}

type Req struct {
	Schoolid int64 `json:"schoolid"`
	Userid   int64 `json:"userid"`
}

//查询学校班级信息（单个）
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_QuerySchoolUserInfo, access_token), req)
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
