package test

import (
	"testing"
	. "util"
)

func ATestLog1(t *testing.T) {
	log := NewLogger("huangsu")
	log.Debugf("debugf %s", "123")
	log.Debugln("debugln %s", "123")
}
