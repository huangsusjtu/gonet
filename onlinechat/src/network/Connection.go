package network

import (
	. "config"
	"net"
	"time"
)

type Connection struct {
	conn             net.Conn
	dataHandler      IRawDataReader
	id               int32
	lastActiveTime   time.Time
	recvBuf          []byte
	outputHeadBuffer *bytes.Buffer
	needClose        bool
}

type IConnectionErrorHandler interface {
	OnConnectionClose(con *Connection)
}

func newConnection(conn net.Conn) *Connection {
	if conn == nil {
		logger.Errorln("conn *net.TCPConn null")
	}
	return &Connection{
		conn:             conn,
		dataHandler:      nil,
		id:               0,
		lastActiveTime:   time.Now(),
		recvBuf:          make([]byte, BUFFER_SIZE),
		needClose:        false,
		outputWriter:     bufio.NewWriter(conn),
		outputHeadBuffer: bytes.NewBuffer([]byte{}),
	}
}

func (self *Connection) WaitAndReadData() {
	var rbyte int
	var err error
	for {
		if self.conn == nil {
			logger.Errorf("self.conn null pointer\n")
			time.Sleep(time.Second)
			return
		}

		rbyte, err = self.conn.Read(self.recvBuf)
		if err != nil {
			logger.Errorln(err.Error())
			self.needClose = true
			return
		}

		self.lastActiveTime = time.Now()
		if rbyte < cap(self.recvBuf) && rbyte > 0 {
			// . . .
			//logger.Debugf("Connection recv  %x\n", self.recvBuf[:rbyte])
			if self.dataHandler != nil {
				self.dataHandler.OnRawDataAvailable(self.id, self.recvBuf[0:rbyte])
			}
		}
	}
}

func (self *Connection) WriteData(raw []byte) {
	if self.needClose == true {
		logger.Debugln("Connection have err, need to close", self)
		return
	}

	if raw == nil || len(raw) == 0 {
		logger.Errorln("Connection no data to send")
		return
	}

	self.outputHeadBuffer.Reset()

	num, err := self.conn.Write(raw)
	if err != nil {
		self.needClose = true
		logger.Errorln(err.Error(), "send data bytes", num)
		return
	}
	//logger.Debugln("send raw data to client ", raw, num)
}

func (self *Connection) recvThread() {

}
