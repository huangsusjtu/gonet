package controller

import (
	. "dispatch"
	"errors"
	"sync"
	. "util"
)

/**
* It Contains all the controllers
* mHanderMap is map of controllers,  key is action ID, value is controller
 */
type ControllerContainer struct {
	mHanderMap        map[int32]IController
	mResponseConsumer IResponseConsumer
	mLock             sync.Mutex
}

var logger = NewLogger("Controller")
var container = ControllerContainer{mHanderMap: make(map[int32]IController, 128)}

func GetControllerContainer() *ControllerContainer {
	return &container
}

func (cc *ControllerContainer) init() {
	// All controllers are config in Route.go
	initRoutTable()

	logger.Debugln("ControllerContainer init ok")
}

func (cc *ControllerContainer) destroy() {
	logger.Debugln("ControllerContainer destroy ok")
}

func (cc *ControllerContainer) GetController(action int32) IController {
	cc.mLock.Lock()
	defer cc.mLock.Unlock()

	if action < 0 {
		logger.Errorf("action id %d illegal\n", action)
		return nil
	}

	if cc.mHanderMap[action] == nil {
		logger.Warnf("action %d, handler is nil\n", action)
		return nil
	}

	return cc.mHanderMap[action]
}

func (cc *ControllerContainer) register(action int32, controller interface{}) {
	if _, ok := controller.(IController); ok == false {
		panic(errors.New("ControllerContainer register parametor type is not IController"))
	}

	cc.mLock.Lock()
	defer cc.mLock.Unlock()

	if action < 0 {
		logger.Errorf("action id %d illegal\n", action)
		panic(errors.New("ControllerContainer register action is nagtive"))
	}

	if cc.mHanderMap[action] != nil {
		panic(errors.New("ControllerContainer register conflict"))
	}

	cc.mHanderMap[action] = controller.(IController)
}

func (cc *ControllerContainer) SetResponseConsumer(responseConsumer IResponseConsumer) {
	cc.mResponseConsumer = responseConsumer
}

func (cc *ControllerContainer) GetResponseConsumer() IResponseConsumer {
	return cc.mResponseConsumer
}

func InitController() *ControllerContainer {
	GetControllerContainer().init()
	return GetControllerContainer()
}

func DestroyController() {
	GetControllerContainer().destroy()
}
