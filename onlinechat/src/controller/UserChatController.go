package controller

import (
	. "dispatch"
	. "model"
	. "protocol"
)

// handle user send message
type UserChatController struct {
}

func (ucc *UserChatController) Handle(context *Context, req *Request) *Response {
	// get user message
	var chatms *ChatMessage = req.GetMs()
	if chatms == nil {
		return &Response{Code: int32(CODE_PARAM_ERROR)}
	}

	// check user login
	userContext := GetUserContextManager().GetUserContext(chatms.Uid)
	if userContext == nil {
		return &Response{Code: int32(CODE_USER_NOLOGIN), Errmsg: "your need login first"}
	}

	// check roominfo
	var roomInfo = chatms.GetRoom()
	if roomInfo == nil || (roomInfo.ChatId <= 0 && len(roomInfo.Uid) < 2) {
		return &Response{Code: int32(CODE_PARAM_ERROR), Errmsg: "need room info"}
	}

	// try to get the target room
	var room *ChatRoom = nil
	if roomInfo.ChatId > 0 { // the room has a id, check exist
		room = GetChatRoomManager().GetChatRoom(roomInfo.ChatId)
	}

	// if room does not exist, create it
	if room == nil {
		room = NewChatRoom(roomInfo.Uid)
	}

	// construct response
	chatms.Room.ChatId = room.Id
	response := &Response{Code: int32(CODE_USER_MESSAGE), Ms: chatms}
	ucc.broacast(room, response)

	return &Response{Code: int32(CODE_OK)}
}

func (ucc *UserChatController) broacast(room *ChatRoom, response *Response) {
	consumer := GetControllerContainer().GetResponseConsumer()

	for _, uid := range room.Uids {
		context := GetUserContextManager().GetUserContext(uid)
		if context == nil { // offline
			//

		} else { // online
			consumer.OnResponseAvailable(context.GetConnection(), response)
		}
	}
}
