package test

import (
	"testing"
	"time"
	. "util"
)

func ATestNewUuid1(t *testing.T) {
	logger := NewLogger("UUID")
	logger.Infoln(NewUuid())
	logger.Infoln(NewUuid())
	logger.Infoln(NewUuid())
	logger.Infoln(NewUuid())
	logger.Infoln(NewUuid())
	time.Sleep(time.Millisecond * 1)
	logger.Infoln(NewUuid())
}
