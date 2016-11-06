package network

// 普通封包
type Packet struct {
	MsgLen uint16 // 消息长度
	Data   []byte
}
