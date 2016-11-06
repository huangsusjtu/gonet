package network

import (
	"errors"
	"fmt"
	"net"
)

type TcpServer struct {
	serverAddr  string
	listener    *net.TCPListener
	connmanager *ConnectionManager
	running     bool
}

func NewTcpServer(addr string, cm *ConnectionManager) *TcpServer {
	server := &TcpServer{
		serverAddr:  addr,
		listener:    nil,
		connmanager: cm,
		running:     false,
	}

	return server
}

func (server *TcpServer) init() (err error) {
	addr, err := net.ResolveTCPAddr("tcp4", server.serverAddr)
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}
	server.listener = listener
	return nil
}

func (server *TcpServer) destroy() (err error) {
	return server.listener.Close()
}

func (server *TcpServer) run() (err error) {
	if server.listener == nil {
		return errors.New("tcpserver not inited")
	}

	server.running = true
	for {
		if server.running == false {
			return nil
		}

		conn, err := server.listener.AcceptTCP()
		if err != nil {
			continue
		}

		conn.SetNoDelay(true)
		conn.SetKeepAlive(true)
		//
		server.connmanager.Put(newConnection(conn))
	}

	return nil
}

func (server *TcpServer) dump() string {
	return fmt.Sprintf("Addr:%s, running:%d", server.serverAddr, server.running)
}
