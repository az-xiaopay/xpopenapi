package add_school_grade

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
type Gradeinfo struct {
	Gradeid   int    `json:"gradeid"`
	Gradename string `json:"gradename"`
}
type Surefire struct {
	Gradeinfo Gradeinfo `json:"gradeinfo"`
	Idx       int       `json:"idx"`
}

type Req struct {
	Schoolid  int64       `json:"schoolid"`
	Opengrade []Opengrade `json:"opengrade"`
}
type Opengrade struct {
	Gradename string `json:"gradename"`
}

//录入学校年级信息
//该接口用于校智付开放平台应用添加当前已授权学校年级信息，支持批量添加
//get_access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_AddSchoolGrade, access_token), req)
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
