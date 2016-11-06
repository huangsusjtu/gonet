package protocol

import (
	proto "github.com/golang/protobuf/proto"
	. "util"
)

var logger = NewLogger("ProtocolStack")

// Serial Request and Response
type IStack interface {
	PackResponse(response *Response) ([]byte, error)
	UnpackResponse(raw []byte) (*Response, error)
	PackRequest(request *Request) ([]byte, error)
	UnpackRequest(raw []byte) (*Request, error)
}

type PBStack struct {
}

func (ps *PBStack) PackResponse(response *Response) ([]byte, error) {
	// . . .
	data, err := proto.Marshal(response)
	return data, err
}

func (ps *PBStack) PackRequest(request *Request) ([]byte, error) {
	// . . .
	data, err := proto.Marshal(request)
	return data, err
}

func (ps *PBStack) UnpackResponse(raw []byte) (*Response, error) {
	response := &Response{}
	err := proto.Unmarshal(raw, response)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return response, nil
}

func (ps *PBStack) UnpackRequest(raw []byte) (*Request, error) {
	request := &Request{}
	err := proto.Unmarshal(raw, request)
	if err != nil {
		logger.Errorln(err)
		return nil, err
	}
	return request, nil
}
