package dispatch

import (
	. "protocol"
)

// filter layer for request and response
// return false means drop request
type ILayer interface {
	OnRequestAvailable(req *Request) bool
	OnResponseAvailable(rep *Response) bool
}

type LayerManager struct {
	mLayers []ILayer
}

func NewLayerManager() *LayerManager {
	//layers := make([]ILayer)
	//layers = append(layers)
	layerManager := &LayerManager{}

	return layerManager
}

// pre handle the request,  return false means drop request
func (m *LayerManager) OnRequestAvailable(req *Request) bool {
	if m.mLayers == nil || len(m.mLayers) == 0 {
		return true
	}

	for _, layer := range m.mLayers {
		if layer.OnRequestAvailable(req) == false {
			return false
		}
	}
	return true
}

// pre handle the response,  return false means drop response
func (m *LayerManager) OnResponseAvailable(rep *Response) bool {
	if m.mLayers == nil || len(m.mLayers) == 0 {
		return true
	}

	for len := len(m.mLayers); len > 0; len-- {
		if m.mLayers[len-1].OnResponseAvailable(rep) == false {
			return false
		}
	}
	return true
}
