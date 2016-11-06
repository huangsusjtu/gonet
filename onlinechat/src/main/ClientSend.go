package main

import (
	. "config"
	. "controller"
	"flag"
	"net"
	. "protocol"
	"time"
	. "util"
)

var logger = NewLogger("client")

func makeLoginRequest(name string, password string) *Request {
	userinfo1 := &UserReqInfo{Type: LOGIN_NAME, Username: name, Password: HashMd5(password)}
	req := &Request{Action: int32(ACTION_USER_LOGIN), Userinfo: userinfo1}
	return req
}

func makeChatRequest() *Request {
	ms := &ChatMessage{Uid: 6189453735362560, Room: &ChatMessage_ChatRoom{Uid: []int64{6190724995350528, 6189453735362560, 6190725100208129, 6190725196677122}}, Content: "hello"}
	req := &Request{Action: int32(ACTION_CHAT_MMS), Ms: ms}
	return req
}

func loginAndSend(name string, password string) {
	conn, err := net.Dial("tcp", SERVER_ADDR)
	if err != nil {
		logger.Errorln(err)
		return
	}
	defer conn.Close()

	stack := &PBStack{}

	data, err := stack.PackRequest(makeLoginRequest(name, password))
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

	for {
		time.Sleep(2 * time.Second)
		data, err := stack.PackRequest(makeChatRequest())
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
	}
}

func main() {
	name := flag.String("user", "", "input your name")
	password := flag.String("p", "1", "input your password")
	flag.Parse()
	loginAndSend(*name, *password)
}
