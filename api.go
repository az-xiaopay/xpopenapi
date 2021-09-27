package xiaopay_openapi

import (
	"xpopenapi/xiaopay/xpapi/add_school_grade"
	"xpopenapi/xiaopay/xptool"
)

func (a openApi) AddSchoolGrade(req *add_school_grade.Req) (*add_school_grade.Resp, error) {
	at, err := xptool.GetAccessToken()
	if err != nil {
		return nil, err
	}
	return add_school_grade.Do(at, req)
}
