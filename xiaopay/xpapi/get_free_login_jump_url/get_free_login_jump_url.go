package get_free_login_jump_url

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
	Schoolid   int64 `json:"schoolid"`   //学校ID
	Userid     int64 `json:"userid"`     //平台用户id, 驻校为user_agent的ID  卡务或学校管理员为user_school的ID
	Menuno     int64 `json:"menuno"`     //菜单编号   1 消费账户 2消费数据报表 3 消费系统
	Usertypept int64 `json:"usertypept"` //用户类型    1 驻校 2卡务 3 消费管理员 （学校管理员）
}

//获取免登录跳转地址
//get_access_token 接口授权令牌
func Do(access_token string, req *Req) (*Resp, error) {
	res, err := xpnet.HttpPost(fmt.Sprintf(xiaopay.API_FreeLoginJump, access_token), req)
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
