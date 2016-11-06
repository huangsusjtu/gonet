package network

import (
	"testing"
	. "util"
)

type Listener struct {
	I int
}

func (l *Listener) OnConnectionAdd(c *Connection) {
	Debugf("listener OnConnectionAdd")
}
func (l *Listener) OnConnectionRemove(c *Connection) {
	Debugf("listener OnConnectionRemove")
}

func TestConnectionManager1(t *testing.T) {
	cm := GetConnectionManager()

	var err error
	var listen = &Listener{10}
	err = cm.RegisterListener(listen)
	if err != nil {
		Debugln(err.Error())
	}

	//cm.Put(NewConnection(nil))

	var listen1 = &Listener{1}
	var listen2 = &Listener{2}
	var listen3 = &Listener{3}
	err = cm.RegisterListener(listen1)
	if err != nil {
		Debugln(err.Error())
	}
	err = cm.RegisterListener(listen2)
	if err != nil {
		Debugln(err.Error())
	}
	err = cm.RegisterListener(listen3)
	if err != nil {
		Debugln(err.Error())
	}
	cm.Put(newConnection(nil))

	err = cm.UnregisterListener(listen2)
	if err != nil {
		Debugln(err.Error())
	}

	cm.Put(newConnection(nil))
}
