package network

import (
	. "config"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
)

type TlsServer struct {
	serverAddr  string
	listener    net.Listener
	connmanager *ConnectionManager
	running     bool
}

func NewTlsServer(addr string, cm *ConnectionManager) *TlsServer {
	server := &TlsServer{
		serverAddr:  addr,
		listener:    nil,
		connmanager: cm,
		running:     false,
	}

	return server
}

func (server *TlsServer) init() (err error) {

	cert, err := tls.LoadX509KeyPair(TLS_PUB_FILE, TLS_RPI_FILE)
	if err != nil {
		logger.Errorf("Error loading certificate. %s \n", err.Error())
		return err
	}

	tlsCfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	listener, err := tls.Listen("tcp4", server.serverAddr, tlsCfg)
	if err != nil {
		logger.Errorln(err)
	}

	server.listener = listener
	return nil
}

func (server *TlsServer) destroy() (err error) {
	return server.listener.Close()
}

func (server *TlsServer) run() (err error) {
	if server.listener == nil {
		return errors.New("tcpserver not inited")
	}

	server.running = true
	for {
		if server.running == false {
			return nil
		}

		conn, err := server.listener.Accept()
		if err != nil {
			continue
		}

		//
		server.connmanager.Put(newConnection(conn))
	}

	return nil
}

func (server *TlsServer) dump() string {
	return fmt.Sprintf("Addr:%s, running:%d", server.serverAddr, server.running)
}
