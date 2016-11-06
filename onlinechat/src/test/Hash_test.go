package test

import (
	"testing"
	. "util"
)

func ATestHash1(t *testing.T) {
	logger := NewLogger("huangsu")
	logger.Debugf("debugf %x\n", HashMd5("123"))
	logger.Debugf("debugln %x\n", HashMd5("123"))
	logger.Debugf("debugln %x\n", HashMd5("123456789"))
	logger.Debugf("debugln %x\n", HashMd5("123456789"))
}
