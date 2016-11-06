package controller

import (
	"model"
	"persistence"
	. "protocol"
	"testing"
)

func ATestUserRegController1(t *testing.T) {
	persistence.GetSqlConPool().Init()

	model.InitModel()
	InitController()

	c := &UserRegController{}

	// check exist
	userinfo1 := &UserReqInfo{Type: CHECK_USER_REGISTED, Username: "huangsu", Email: "su.huang@qq.com", Phone: "13166357974", Desc: "...", Password: "..."}
	req1 := &Request{Action: int32(ACTION_USER_REG), Userinfo: userinfo1}
	rep1 := c.Handle(nil, req1)
	logger.Debugf("%v %d\n", rep1, rep1.Code)

	// register
	userinfo2 := &UserReqInfo{Type: REGISTER, Username: "huangsu", Email: "su.huang@qq.com", Phone: "13166357974", Desc: "...", Password: "123456789"}
	req2 := &Request{Action: int32(ACTION_USER_REG), Userinfo: userinfo2}
	rep2 := c.Handle(nil, req2)
	logger.Debugf("%v %d\n", rep2, rep2.Code)

	// check exist again
	rep1 = c.Handle(nil, req1)
	logger.Debugf("%v %d\n", rep1, rep1.Code)

	// register again
	rep2 = c.Handle(nil, req2)
	logger.Debugf("%v %d\n", rep2, rep2.Code)

	// register
	userinfo3 := &UserReqInfo{Type: REGISTER, Username: "huangsu1", Email: "su.huang1@qq.com", Phone: "131663579741", Desc: "...", Password: "123456789"}
	req3 := &Request{Action: int32(ACTION_USER_REG), Userinfo: userinfo3}
	rep3 := c.Handle(nil, req3)
	logger.Debugf("%v %d\n", rep3, rep3.Code)

	c.destroy()

	DestroyController()
	model.DestroyModel()
	persistence.GetSqlConPool().Destroy()
}
