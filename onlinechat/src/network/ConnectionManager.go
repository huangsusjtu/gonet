package network

import (
	"errors"
	"sync"
	. "util"
)

type ConnectionManager struct {
	connMap     map[int32]*Connection
	allocConnID int32
	listeners   []IConnectionListener
	handler     IRawDataReader
	lock        sync.Mutex
}

type IConnectionListener interface {
	OnConnectionAdd(c *Connection)
	OnConnectionRemove(c *Connection)
}

var cm = ConnectionManager{connMap: make(map[int32]*Connection, 1024), allocConnID: 0, listeners: make([]IConnectionListener, 0)}
var logger = NewLogger("network")

func GetConnectionManager() *ConnectionManager {
	return &cm
}

func (cm *ConnectionManager) SetDataHandler(handler IRawDataReader) {
	cm.handler = handler
}

func (cm *ConnectionManager) Put(connection *Connection) (id int32) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	var exist bool
	for exist = false; exist == true; cm.allocConnID++ {
		_, exist = cm.connMap[cm.allocConnID]
		logger.Warnf("current Id:%s is used\n", cm.allocConnID)
	}

	cm.connMap[cm.allocConnID] = connection
	id = cm.allocConnID
	connection.id = id
	connection.dataHandler = cm.handler

	cm.allocConnID++

	for _, listen := range cm.listeners {
		(listen).OnConnectionAdd(connection)
	}

	go connection.WaitAndReadData()
	return id
}

func (cm *ConnectionManager) Get(id int32) *Connection {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	conn, exist := cm.connMap[id]
	if exist == true {
		return conn
	}
	return nil
}

func (cm *ConnectionManager) Remove(id int32) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	conn, exist := cm.connMap[id]
	if exist == false {
		return
	}

	for _, listen := range cm.listeners {
		(listen).OnConnectionRemove(conn)
	}
	delete(cm.connMap, id)
	Debugf("remove %d connection\n", id)
}

func (cm *ConnectionManager) RegisterListener(l interface{}) error {

	if _, ok := l.(IConnectionListener); ok == false {
		return errors.New("parametor type is not IConnectionListener")
	}

	cm.lock.Lock()
	defer cm.lock.Unlock()

	for _, listen := range cm.listeners {
		if l == listen {
			return errors.New("listener exist")
		}
	}

	cm.listeners = append(cm.listeners[:], l.(IConnectionListener))
	return nil
}

func (cm *ConnectionManager) UnregisterListener(l interface{}) error {
	if _, ok := l.(IConnectionListener); ok == false {
		return errors.New("parametor type is not IConnectionListener")
	}

	cm.lock.Lock()
	defer cm.lock.Unlock()

	for i, listen := range cm.listeners {
		if l != listen {
			continue
		}

		if len(cm.listeners) == 1 {
			cm.listeners = make([]IConnectionListener, 0)
			return nil
		}

		listeners := make([]IConnectionListener, len(cm.listeners)-1)
		if i == len(cm.listeners)-1 {
			copy(listeners, cm.listeners[0:i])
		} else if i == 0 {
			copy(listeners, cm.listeners[i+1:])
		} else {
			copy(listeners[0:i], cm.listeners[0:i])
			copy(listeners[i:], cm.listeners[i+1:])
		}
		cm.listeners = listeners
		return nil
	}
	return errors.New("this listener is not registered")
}
