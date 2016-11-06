package model

import (
	. "util"
)

var logger = NewLogger("Model")

func InitModel() {
	// init each model
	initUserModel()
}

func DestroyModel() {
	// destroy each model
	destroyUserModel()

}
