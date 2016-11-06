package controller

import (
	"model"
	"persistence"
	. "protocol"
	"testing"
	"util"
)

func TestUserLoginController1(t *testing.T) {
	persistence.GetSqlConPool().Init()

	model.InitModel()
	InitController()

	c := &UserLoginController{}

	// login
	userinfo1 := &UserReqInfo{Type: LOGIN_PHONE, Username: "huangsu", Email: "su.huang@qq.com", Phone: "13166357974", Desc: "...", Password: util.HashMd5("123456789")}
	req1 := &Request{Action: int32(ACTION_USER_LOGIN), Userinfo: userinfo1}
	rep1 := c.Handle(nil, req1)
	logger.Debugf("%v %d\n", rep1, rep1.Code)

	c.destroy()

	DestroyController()
	model.DestroyModel()
	persistence.GetSqlConPool().Destroy()
}
