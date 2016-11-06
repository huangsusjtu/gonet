package session

import (
	"config"
	"strings"
	"util"
)

type ISessionObject interface {
}

type ISession interface {
	putSession(o interface{}) (sessionID int64)
	getSession(id int64) interface{}
}

type SessionManager struct {
	mSession  ISession
	mTotalNum int64
}

func (sm *SessionManager) Init() {
	if strings.EqualFold(config.SESSION_TYPE, "memory") {

	}
}

func (sm *SessionManager) Destroy() {

}
