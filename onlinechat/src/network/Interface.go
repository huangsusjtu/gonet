package network

// Interface used by Dispatcher and Network level
// Dispatcher level use it to send data
type IRawDataWriter interface {
	WriteData(raw []byte)
}

// Dispatcher level implement it to read data
type IRawDataReader interface {
	OnRawDataAvailable(connid int32, raw []byte)
}
