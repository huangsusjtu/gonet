package model

import (
	"sync"
)

type ChatRoom struct {
	Id   int64
	Uids []int64
}

type ChatRoomManager struct {
	mRoomMap map[int64]*ChatRoom
	mLock    sync.Mutex
	mNextId  int64
}

var chatRoomManager = &ChatRoomManager{mRoomMap: make(map[int64]*ChatRoom)}

func GetChatRoomManager() *ChatRoomManager {
	return chatRoomManager
}

func (crm *ChatRoomManager) allocID() int64 {
	crm.mLock.Lock()
	defer crm.mLock.Unlock()
	crm.mNextId++
	return crm.mNextId
}

func (crm *ChatRoomManager) Add(room *ChatRoom) {
	if room == nil || room.Id < 0 || len(room.Uids) < 2 {
		logger.Debugln("Invalid ChatRoom ", room)
		return
	}

	crm.mLock.Lock()
	defer crm.mLock.Unlock()
	if crm.mRoomMap[room.Id] != nil {
		logger.Errorln("ChatRoomManager Add Room exist ", room)
		return
	}
	crm.mRoomMap[room.Id] = room
}

func (crm *ChatRoomManager) Remove(id int64) {
	if id < 0 {
		return
	}

	crm.mLock.Lock()
	defer crm.mLock.Unlock()
	delete(crm.mRoomMap, id)
}

func (crm *ChatRoomManager) GetChatRoom(id int64) *ChatRoom {
	crm.mLock.Lock()
	defer crm.mLock.Unlock()

	return crm.mRoomMap[id]
}

func NewChatRoom(uids []int64) *ChatRoom {
	room := &ChatRoom{Id: GetChatRoomManager().allocID()}
	room.Uids = make([]int64, len(uids))
	copy(room.Uids, uids)
	GetChatRoomManager().Add(room)
	return room
}
