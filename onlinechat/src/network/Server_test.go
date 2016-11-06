package network

import (
	. "config"
	"net"
	"testing"
	"time"
	. "util"
)

func initClient() {
	conn, err := net.Dial("tcp", SERVER_ADDR)
	if err != nil {
		// handle error

		Errorf("new tcp client err %s\n", err.Error())
		return
	}
	conn.Write([]byte("1234"))

	conn.Close()
}

func TestServer1(t *testing.T) {
	cm := GetConnectionManager()
	sm := GetDefaultServerManager()
	sm.Add(NewTcpServer(SERVER_ADDR, cm))
	sm.RunAllServer()

	for i := 0; i < 1; i++ {
		initClient()
	}
	time.Sleep(5 * time.Second)
}
