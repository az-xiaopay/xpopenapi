package xiaopay

import "xpopenapi/common/config"

var (
	API_application_token       = config.Base.XPServerUrl + "/openapi/application/token"
	API_AddSchoolGrade          = config.Base.XPServerUrl + "/openapi/basic_business/AddSchoolGrade?token=%s"
	API_AddSchoolClass          = config.Base.XPServerUrl + "/openapi/basic_business/AddSchoolClass?token=%s"
	API_AddSchoolUser           = config.Base.XPServerUrl + "/openapi/basic_business/AddSchoolUser?token=%s"
	API_EditSchoolGrade         = config.Base.XPServerUrl + "/openapi/basic_business/EditSchoolGrade?token=%s"
	API_EditSchoolClass         = config.Base.XPServerUrl + "/openapi/basic_business/EditSchoolClass?token=%s"
	API_EditSchoolUser          = config.Base.XPServerUrl + "/openapi/basic_business/EditSchoolUser?token=%s"
	API_DelSchoolGrade          = config.Base.XPServerUrl + "/openapi/basic_business/DelSchoolGrade?token=%s"
	API_DelSchoolClass          = config.Base.XPServerUrl + "/openapi/basic_business/DelSchoolClass?token=%s"
	API_DelSchoolUser           = config.Base.XPServerUrl + "/openapi/basic_business/DelSchoolUser?token=%s"
	API_SendSMSAuthCode         = config.Base.XPServerUrl + "/openapi/basic_business/SendSMSAuthCode?token=%s"
	API_CheckSMSCode            = config.Base.XPServerUrl + "/openapi/basic_business/CheckSMSCode?token=%s"
	API_FreeLoginJump           = config.Base.XPServerUrl + "/openapi/basic_business/FreeLoginJump?token=%s"
	API_WeappFreeLoginJump      = config.Base.XPServerUrl + "/openapi/basic_business/WeappFreeLoginJump?token=%s"
	API_CreateUserSchool        = config.Base.XPServerUrl + "/openapi/basic_business/CreateUserSchool?token=%s"
	API_QueryOneSchoolClassInfo = config.Base.XPServerUrl + "/openapi/basic_business/QueryOneSchoolClassInfo?token=%s"
	API_QueryOneSchoolGradeInfo = config.Base.XPServerUrl + "/openapi/basic_business/QueryOneSchoolGradeInfo?token=%s"
	API_QuerySchoolUserInfo     = config.Base.XPServerUrl + "/openapi/basic_business/QuerySchoolUserInfo?token=%s"
	API_AddOpenAppSchool        = config.Base.XPServerUrl + "/openapi/basic_business/AddOpenAppSchool?token=%s"
	API_AddOpenAppInterface     = config.Base.XPServerUrl + "/openapi/basic_business/AddOpenAppInterface?token=%s"
	API_QueryUserSchool         = config.Base.XPServerUrl + "/openapi/basic_business/QueryUserSchool?token=%s"
)
