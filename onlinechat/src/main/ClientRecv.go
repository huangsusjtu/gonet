package main

import (
	. "config"
	. "controller"
	"flag"
	"net"
	. "protocol"
	. "util"
)

var logger = NewLogger("client")

func makeLoginRequest(name string, password string) *Request {
	userinfo1 := &UserReqInfo{Type: LOGIN_NAME, Username: name, Password: HashMd5(password)}
	req := &Request{Action: int32(ACTION_USER_LOGIN), Userinfo: userinfo1}
	return req
}

func loginAndWait(name string, password string) {
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

func main() {
	name := flag.String("user", "hehe", "input your name")
	password := flag.String("p", "1", "input your password")
	flag.Parse()
	logger.Debugln(*name, *password)
	loginAndWait(*name, *password)
}
