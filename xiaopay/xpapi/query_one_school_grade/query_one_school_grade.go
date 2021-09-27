package query_one_school_grade

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
	Gradeid   int64  `json:"gradeid"`
	Gradename string `json:"gradename"`
}

type Req struct {
	Schoolid int64 `json:"schoolid"`
	Gradeid  int64 `json:"gradeid"`
}

//查询学校班级信息（单个）
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_QueryOneSchoolGradeInfo, access_token), req)
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
