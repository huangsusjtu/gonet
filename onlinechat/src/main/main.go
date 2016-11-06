package main

import (
	"controller"
	"dispatch"
	"model"
	"network"
	"persistence"
	"time"
)

func main() {
	// Database init
	persistence.GetSqlConPool().Init()
	// Model init
	model.InitModel()
	// Controller init,  return controller container
	container := controller.InitController()

	// Dispatch is used for dispatch request and response
	var dispatcher = dispatch.InitDispatch(container)
	container.SetResponseConsumer(dispatcher)

	// Server init
	network.InitServer()
	network.GetConnectionManager().SetDataHandler(dispatcher) // data handler will process data from connection

	for {
		time.Sleep(5 * time.Second)
	}

	controller.DestroyController()
	model.DestroyModel()
	persistence.GetSqlConPool().Destroy()
}
