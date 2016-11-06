package network

import (
	"config"
	"errors"
	"sync"
	. "util"
)

type IServer interface {
	init() (err error)
	destroy() (err error)

	run() (err error)
	dump() string
}

type ServerManager struct {
	serverlist []IServer
	lock       sync.Mutex
}

func newServerManager() *ServerManager {
	return &ServerManager{
		serverlist: make([]IServer, 0),
	}
}

var serverManager = newServerManager()

func GetDefaultServerManager() *ServerManager {
	return serverManager
}

func (sm *ServerManager) Add(server interface{}) error {
	if _, ok := server.(IServer); ok == false {
		return errors.New("parametor type is not IServer")
	}

	sm.lock.Lock()
	defer sm.lock.Unlock()

	sm.serverlist = append(sm.serverlist, server.(IServer))
	return nil
}

func (sm *ServerManager) Dump() string {
	return ""
}

func (sm *ServerManager) RunAllServer() {
	var serverlist []IServer
	sm.lock.Lock()
	serverlist = make([]IServer, len(sm.serverlist))
	copy(serverlist, sm.serverlist)
	sm.lock.Unlock()

	for _, server := range serverlist {
		err := server.init()
		if err != nil {
			Errorln(err.Error())
			continue
		}
		go server.run()
	}
}

func InitServer() *ConnectionManager {
	cm := GetConnectionManager()    // all connections
	sm := GetDefaultServerManager() //all server
	sm.Add(NewTcpServer(config.SERVER_ADDR, cm))
	sm.RunAllServer() // start each server
	logger.Debugln("Init net server ok")
	return cm
}
