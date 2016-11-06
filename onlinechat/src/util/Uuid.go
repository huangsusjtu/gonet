package util

import (
	"sync"
	"time"
)

var lock sync.Mutex
var userSeq int64 = 0
var machineID int64 = 0 << 12

// create a unique UID
func NewUuid() int64 {
	lock.Lock()
	seq := userSeq & 0x00000fff
	userSeq = seq + 1
	lock.Unlock()

	t := time.Now().Unix()
	t = (t & 0x0000000001ffffffffff) << 22

	return (t | seq | machineID)
}

var roomSeq int64 = 0

// create a unique UID
func NewRoomId() int64 {
	lock.Lock()
	seq := roomSeq & 0x00000fff
	roomSeq = seq + 1
	lock.Unlock()

	t := time.Now().Unix()
	t = (t & 0x0000000001ffffffffff) << 22

	return (t | seq | machineID)
}
