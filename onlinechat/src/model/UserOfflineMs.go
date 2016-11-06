package model

import (
	"sync"
	"time"
)

type PendingMessage struct {
	mUid    int64
	mRoom   *ChatRoom
	Content string
	t       time.Time
}

// for offline user, store message here
type PendingMessageManager struct {
	mUserMap map[int64]PendingMessage
	lock     sync.Mutex
}
