package controller

import (
	. "dispatch"
	. "model"
	. "protocol"
	"strings"
	"util"
)

type UserRegController struct {
}

/*
message UserRegInfo {
	int32 type = 1;
	string username = 2;
	string email = 3;
	string phone = 4;
	string password = 5;
}
*/

// register type
const (
	CHECK_USER_REGISTED = 100
	REGISTER            = 101
)

func (urc *UserRegController) Handle(context *Context, req *Request) *Response {
	info := req.GetUserinfo()
	if info == nil {
		return &Response{Code: int32(CODE_PARAM_ERROR)}
	}

	var user = UserInfo{Name: info.Username, Email: info.Email, Phone: info.Phone, Password: info.Password, Description: info.Desc}
	var exist bool
	switch info.Type {
	case CHECK_USER_REGISTED:
		if exist = user.CheckUserExist(); exist == true {
			return &Response{Code: int32(CODE_USER_EXIST)}
		}
		return &Response{Code: int32(CODE_USER_NOEXIST)}
	case REGISTER:
		// check param
		if strings.EqualFold(info.Username, "") && strings.EqualFold(info.Email, "") && strings.EqualFold(info.Phone, "") {
			return &Response{Code: int32(CODE_PARAM_ERROR)}
		}
		if strings.EqualFold(user.Password, "") {
			return &Response{Code: int32(CODE_PARAM_ERROR)}
		}

		user.Password = util.HashMd5(user.Password)

		// check user exist
		if exist = user.CheckUserExist(); exist == true {
			return &Response{Code: int32(CODE_USER_EXIST)}
		}

		// register
		ok := user.Register()
		if ok != true {
			return &Response{Code: int32(CODE_SERVER_ERROR)}
		}

		// check register ok
		if exist = user.CheckUserExist(); exist == false {
			return &Response{Code: int32(CODE_SERVER_ERROR)}
		}
		return &Response{Code: int32(CODE_OK)}
	default:
		return &Response{Code: int32(CODE_PARAM_ERROR)}
	}
	return &Response{Code: int32(CODE_OK)}
}
