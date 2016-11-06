package test

import (
	. "config"
	. "controller"
	"net"
	. "protocol"
	"testing"
	"time"
	. "util"
)

var logger = NewLogger("client")

func makeRegisterRequest() *Request {
	userinfo := &UserReqInfo{Type: REGISTER, Username: "huangsu2", Email: "su.huang2@qq.com", Phone: "13166357972", Desc: "2", Password: "1"}
	req := &Request{Action: int32(ACTION_USER_REG), Userinfo: userinfo}
	return req
}

func makeLoginRequest1() *Request {
	userinfo1 := &UserReqInfo{Type: LOGIN_NAME, Username: "huangsu", Password: HashMd5("1")}
	req := &Request{Action: int32(ACTION_USER_LOGIN), Userinfo: userinfo1}
	return req
}

func makeLoginRequest2() *Request {
	userinfo1 := &UserReqInfo{Type: LOGIN_NAME, Username: "huangsu1", Password: HashMd5("1")}
	req := &Request{Action: int32(ACTION_USER_LOGIN), Userinfo: userinfo1}
	return req
}

func makeLoginRequest3() *Request {
	userinfo1 := &UserReqInfo{Type: LOGIN_NAME, Username: "huangsu2", Password: HashMd5("1")}
	req := &Request{Action: int32(ACTION_USER_LOGIN), Userinfo: userinfo1}
	return req
}

func makeLoginRequest4() *Request {
	userinfo1 := &UserReqInfo{Type: LOGIN_NAME, Username: "huangsu3", Password: HashMd5("1")}
	req := &Request{Action: int32(ACTION_USER_LOGIN), Userinfo: userinfo1}
	return req
}

func makeChatRequest() *Request {
	ms := &ChatMessage{Uid: 6189453735362560, Room: &ChatMessage_ChatRoom{Uid: []int64{6189453735362560, 6190724995350528, 6190725100208129, 6190725196677122}}, Content: "hello"}
	req := &Request{Action: int32(ACTION_CHAT_MMS), Ms: ms}
	return req
}

// login and wait for message
func client1(req1 *Request) {
	conn, err := net.Dial("tcp", SERVER_ADDR)
	if err != nil {
		logger.Errorln(err)
		return
	}
	defer conn.Close()

	stack := &PBStack{}

	data, err := stack.PackRequest(req1)
	if err != nil {
		logger.Errorln(err)
		return
	}

	conn.Write(data)
	logger.Debugln("send ok")
	var buffer = make([]byte, 1024)

	for {
		var n int
		if n, err = conn.Read(buffer); err != nil {
			logger.Errorln(err)
			return
		}
		logger.Debugln("recv response ok")
		response, err := stack.UnpackResponse(buffer[:n])
		if err != nil {
			logger.Errorln(err)
			return
		}
		logger.Debugln(response)
	}
}

// client2()
func client2() {
	conn, err := net.Dial("tcp", SERVER_ADDR)
	if err != nil {
		logger.Errorln(err)
		return
	}
	defer conn.Close()

	stack := &PBStack{}

	// login
	data, err := stack.PackRequest(makeLoginRequest4())
	if err != nil {
		logger.Errorln(err)
		return
	}

	conn.Write(data)
	logger.Debugln("send ok")
	var buffer = make([]byte, 1024)

	var n int
	if n, err = conn.Read(buffer); err != nil {
		logger.Errorln(err)
		return
	}
	logger.Debugln("recv response ok")
	response, err := stack.UnpackResponse(buffer[:n])
	if err != nil {
		logger.Errorln(err)
		return
	}
	logger.Debugln(response)

	// send message
	data, err = stack.PackRequest(makeChatRequest())
	if err != nil {
		logger.Errorln(err)
		return
	}

	conn.Write(data)
	logger.Debugln("send ok")
	if n, err = conn.Read(buffer); err != nil {
		logger.Errorln(err)
		return
	}
	logger.Debugln("recv response ok")
	response, err = stack.UnpackResponse(buffer[:n])
	if err != nil {
		logger.Errorln(err)
		return
	}
	logger.Debugln(response)
}

func TestClient(t *testing.T) {
	go client1(makeLoginRequest1())
	go client1(makeLoginRequest2())
	go client1(makeLoginRequest3())
	time.Sleep(5 * time.Second)

	client2()

	time.Sleep(5 * time.Second)
	//client(makeLoginRequest())
}
