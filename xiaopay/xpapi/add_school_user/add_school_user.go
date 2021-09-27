package add_school_user

import (
	"encoding/json"
	"fmt"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
)

type Resp struct {
	xiaopay.RespBase
	Data Data `json:"Data"`
}
type Data struct {
	Fails    []Fails    `json:"fails"`
	Surefire []Surefire `json:"surefire"`
}

type Fails struct {
	Idx     int    `json:"idx"`
	Message string `json:"message"`
}
type Userinfo struct {
	Cardnumber           string `json:"cardnumber"`
	Classid              int    `json:"classid"`
	Classname            string `json:"classname"`
	Gender               int    `json:"gender"`
	Gradeid              int    `json:"gradeid"`
	Gradename            string `json:"gradename"`
	Identificationnumber string `json:"identificationnumber"`
	Identificationtype   string `json:"identificationtype"`
	Phonenumber          string `json:"phonenumber"`
	Userid               int    `json:"userid"`
	Username             string `json:"username"`
	Usertype             string `json:"usertype"`
}
type Surefire struct {
	Idx      int      `json:"idx"`
	Userinfo Userinfo `json:"userinfo"`
}

//============================

type Req struct {
	Schoolid int64      `json:"schoolid"`
	Openuser []Openuser `json:"openuser"`
}
type Openuser struct {
	Gradeid              int64  `json:"gradeid"`
	Classid              int64  `json:"classid"`
	Usertype             string `json:"usertype"`
	Username             string `json:"username"`
	Cardnumber           string `json:"cardnumber"`
	Gender               int    `json:"gender"`
	Phonenumber          string `json:"phonenumber"`
	Identificationtype   string `json:"identificationtype"`
	Identificationnumber string `json:"identificationnumber"`
}

//录入学校用户资料
//该接口用于校智付开放平台应用添加当前已授权学校用户资料信息，支持批量添加
//get_access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_AddSchoolUser, access_token), req)
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
