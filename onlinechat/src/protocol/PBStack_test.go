package protocol

import (
	"testing"
)

func TestPS1(t *testing.T) {

	ps := &PBStack{}

	request := &Request{Action: 1000, Userinfo: &UserReqInfo{Type: 1}}
	data, err := ps.Pack(request)
	if err != nil {
		t.Error(err.Error())
		return
	}

	ps.Unpack(data)
}
