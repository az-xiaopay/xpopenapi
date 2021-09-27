package edit_school_user

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
	Schoolid             int64  `json:"schoolid"`             //学校ID
	Userid               int64  `json:"userid"`               //用户id
	Gradeid              int64  `json:"gradeid"`              //年级ID
	Classid              int64  `json:"classid"`              //班级ID
	Usertype             string `json:"usertype"`             //用户类型 STUDENT:学生 TEACHER:教职工
	Username             string `json:"username"`             //用户姓名
	Cardnumber           string `json:"cardnumber"`           //M1卡序列号
	Gender               int    `json:"gender"`               //性别 0 女 1 男 2 不公开
	Phonenumber          string `json:"phonenumber"`          //家长手机号
	Identificationtype   string `json:"identificationtype"`   //证件类型
	Identificationnumber string `json:"identificationnumber"` //证件号码 IDCARD 身份证 EEP_HK_MACAU（港澳通行证）,PASSPORT_NO（护照）

}

//编辑学校用户信息
//get_access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_EditSchoolUser, access_token), req)
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
