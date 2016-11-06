package controller

import (
	"persistence"
	. "protocol"
	"testing"
)

type controller struct {
	index int
}

func (c *controller) init() {
	logger.Debugln("controller init")
}

func (c *controller) Handle(req *Request) *Response {
	logger.Debugln("test controller", c.index)
	return nil
}

func (c *controller) destroy() {
	logger.Debugln("controller destroy")
}

func ATestController1(t *testing.T) {
	persistence.GetSqlConPool().Init()

	container := GetControllerContainer()
	container.register(10001, &controller{index: 1234})

	container.init()

	c := container.getController(0)
	if c != nil {
		c.Handle(nil, nil)
	}

	container.destroy()
	persistence.GetSqlConPool().Destroy()
}
