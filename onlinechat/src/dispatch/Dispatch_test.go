package dispatch

import (
	"model"
	"persistence"
	. "protocol"
	"testing"
	"time"
)

func ATestRequestDispatch1(t *testing.T) {
	persistence.GetSqlConPool().Init()

	model.InitModel()
	InitController()

	rd := NewDispatcher(nil)
	go rd.DispatchRequest()

	userinfo1 := &UserReqInfo{Type: REGISTER, Username: "huangsu", Email: "su.huang@qq.com", Phone: "13166357974", Desc: "...", Password: "..."}
	req1 := &Request{Action: int32(ACTION_USER_REG), Userinfo: userinfo1}

	rd.OnRequestAvailable(req1)

	userinfo2 := &UserReqInfo{Type: REGISTER, Username: "huangsu2", Email: "su.huang2@qq.com", Phone: "13166357975", Desc: "...", Password: "..."}
	req2 := &Request{Action: int32(ACTION_USER_REG), Userinfo: userinfo2}

	rd.OnRequestAvailable(req2)

	time.Sleep(5 * time.Second)

	DestroyController()
	model.DestroyModel()
	persistence.GetSqlConPool().Destroy()
}
