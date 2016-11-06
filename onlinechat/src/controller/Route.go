package controller

import (
	. "protocol"
)

func initRoutTable() {
	// Add controller
	GetControllerContainer().register(int32(ACTION_USER_REG), &UserRegController{})
	GetControllerContainer().register(int32(ACTION_USER_LOGIN), &UserLoginController{})
	GetControllerContainer().register(int32(ACTION_CHAT_MMS), &UserChatController{})
}
