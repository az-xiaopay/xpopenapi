package xpnotify

type Notify struct {
	NotifyId      string      `json:"notify_id"`
	NotifyType    string      `json:"notify_type"`
	NotifyContent interface{} `json:"notify_content"`
}

type Resp struct {
	Code string `json:"return_code"`
	Msg  string `json:"return_msg"`
}

//重采确认通知
type NotifyReacquire struct {
	Appid    string `json:"appid"`
	Noncestr string `json:"noncestr"`
	Sign     string `json:"sign"`
	Signtype string `json:"signtype"`
	Schoolid int64  `json:"schoolid"`
	Userid   int64  `json:"userid"`
	Phone    string `json:"phone"`
	Stucode  string `json:"stucode"`
	Usertype int    `json:"usertype"`
	Xpappid  string `json:"xpappid"`
	Pathurl  string `json:"pathurl"`
}

//消费结果
type NotifyConsumeRes struct {
	Appid       string `json:"appid"`
	Noncestr    string `json:"noncestr"`
	Sign        string `json:"sign"`
	Signtype    string `json:"signtype"`
	Schoolid    int64  `json:"schoolid"`
	Devicecode  string `json:"devicecode"`
	Userid      int64  `json:"userid"`
	Ordernumber string `json:"ordernumber"`
	Dealtime    string `json:"dealtime"`
	Amount      int64  `json:"amount"`
	Stucode     string `json:"stucode"`
	Phone       string `json:"phone"`
	Usertype    string `json:"usertype"`
	ConsumeDesc string `json:"consumedesc"`
	Wallettype  uint8  `json:"wallettype"`
}

//数据变动通知报文
type NotifyDataChange struct {
	Appid      string `json:"appid"`
	Noncestr   string `json:"noncestr"`
	Sign       string `json:"sign"`
	Signtype   string `json:"signtype"`
	Infotype   string `json:"infotype"` //信息类型 基础数据
	Timestamp  int64  `json:"timestamp"`
	Schoolid   int64  `json:"schoolid"`
	Changetype string `json:"changetype"` //修改类型 新增、编辑、删除
	Xpid       int64  `json:"xpid"`       //校智付id
	Outid      string `json:"outid"`      //外部id
}
