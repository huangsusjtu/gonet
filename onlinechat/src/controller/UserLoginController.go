package controller

import (
	. "dispatch"
	. "model"
	. "protocol"
	"strings"
)

type UserLoginController struct {
}

// login type
const (
	LOGIN_NAME  = 200
	LOGIN_EMAIL = 201
	LOGIN_PHONE = 202
)

func (ulc *UserLoginController) Handle(context *Context, req *Request) *Response {
	info := req.GetUserinfo()
	if info == nil || strings.EqualFold(info.Password, "") {
		return &Response{Code: int32(CODE_PARAM_ERROR)}
	}

	var user = UserInfo{Name: info.Username, Email: info.Email, Phone: info.Phone, Password: info.Password, Description: info.Desc}
	var result bool = false
	var errmsg string
	switch info.Type {
	case LOGIN_NAME:
		result, errmsg = user.AuthByName()
	case LOGIN_EMAIL:
		result, errmsg = user.AuthByEmail()
	case LOGIN_PHONE:
		result, errmsg = user.AuthByPhone()
	default:
		return &Response{Code: int32(CODE_PARAM_ERROR)}
	}

	if result == true {
		// auth ok
		// create online user context, and Add to ontextManager
		GetUserContextManager().Add(NewUserContext(&user, context.GetConnection()))
		logger.Debugln("user login ok", user.Name)
		return &Response{Code: int32(CODE_USER_AUTH_SUCCESS), UserId: user.Uid}
	}
	return &Response{Code: int32(CODE_USER_AUTH_FAIL), Errmsg: errmsg}
}
