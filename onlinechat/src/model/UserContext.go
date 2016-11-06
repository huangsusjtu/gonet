package model

import (
	"config"
	"sync"
	"time"
)

// An online user has a context
type UserContext struct {
	mUser    UserInfo
	mOnline  bool
	mExpired time.Time
	mConnId  int32
}

// create a instance of UserContext
func NewUserContext(user *UserInfo, connid int32) *UserContext {
	context := &UserContext{mUser: *user, mOnline: true, mConnId: connid}
	context.UpdateExpiredTime()
	return context
}

// update the expired time of online user
func (u *UserContext) UpdateExpiredTime() {
	t := time.Now()
	t.Add(time.Duration(config.EXPIRED_TIME_SEC) * time.Second)
	u.mExpired = t
}

func (u *UserContext) GetConnection() int32 {
	return u.mConnId
}

// container of all online users
type UserContextManager struct {
	mUserMap map[int64]*UserContext
	mLock    sync.Mutex
}

// Single instance of ContextManager
var userContextManager = &UserContextManager{mUserMap: make(map[int64]*UserContext, config.DEFAULT_ONLINE_USER)}

// get the default ContextManager
func GetUserContextManager() *UserContextManager {
	return userContextManager
}

// Add a context to ContextManager
func (m *UserContextManager) Add(c *UserContext) {
	if c == nil || c.mUser.Uid == 0 {
		logger.Errorln("UserContextManager Add Param err ", c)
		return
	}

	m.mLock.Lock()
	defer m.mLock.Unlock()
	uid := c.mUser.Uid
	if m.mUserMap[uid] != nil {
		logger.Errorln("UserContextManager Add Context exist ", c)
		return
	}

	m.mUserMap[uid] = c
}

// Remove a context from ContextManager
func (m *UserContextManager) Remove(uid int64) {
	if uid == 0 {
		logger.Errorln("UserContextManager Add Param err ", uid)
		return
	}

	m.mLock.Lock()
	defer m.mLock.Unlock()
	delete(m.mUserMap, uid)
}

func (m *UserContextManager) GetUserContext(uid int64) *UserContext {
	m.mLock.Lock()
	defer m.mLock.Unlock()

	return m.mUserMap[uid]
}
