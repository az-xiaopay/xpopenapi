package edit_school_class

import (
	"encoding/json"
	"fmt"
	"xpopenapi/common/xpnet"
	"xpopenapi/xiaopay"
)

type Resp struct {
	xiaopay.RespBase
	Data string `json:"Data"`
}

type Req struct {
	Schoolid  int64  `json:"schoolid"`  //学校ID
	Gradeid   int64  `json:"gradeid"`   //年级ID
	Classid   int64  `json:"classid"`   //班级ID
	ClassName string `json:"classname"` //班级名称
}

//编辑学校班级信息
//get_access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_EditSchoolClass, access_token), req)
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
