package session

import (
	. "chat/config"
	"chat/corelib/logger"
	. "chat/corelib/network"
	. "chat/types"
	"encoding/json"
	"net"
)

type PacketSession struct {
	id         int
	conn       net.Conn
	isShutdown bool
	sendChan   chan MessagePacket
	recvChan   chan MessagePacket
}

func NewPacketSession(InID int, InConn net.Conn) (*PacketSession, error) {
	session := &PacketSession{
		id:       InID,
		conn:     InConn,
		sendChan: make(chan MessagePacket),
		recvChan: make(chan MessagePacket),
	}

	// PacketEvent Bind!
	go session.startPacketEvent()

	// Receive Start!
	go session.startReceive()

	return session, nil
}

func (self *PacketSession) Send(InPacket MessagePacket) error {
	self.sendChan <- InPacket
	return nil
}

func (self *PacketSession) Shutdown() {
	if self.conn != nil {
		disconnectErr := self.conn.Close()
		if disconnectErr != nil {
			logger.Error(disconnectErr.Error())
		}
	}
}

func (self *PacketSession) startPacketEvent() {
	for {
		select {
		case sendPkt := <-self.sendChan:
			sendBytes, marshalErr := json.Marshal(sendPkt)
			if marshalErr != nil {
				logger.Error(marshalErr.Error())
				continue
			}

			self.conn.Write(sendBytes)
		case recvPkt := <-self.recvChan:
			logger.Text(recvPkt.Message)
		}
	}
}

func (self *PacketSession) startReceive() {

	recvBuffer := make([]byte, PACKET_MAX_BUFFER)
	recvPacket := MessagePacket{0, ""}

	var readSize int = 0
	var packetSize int32 = 0
	var messageData string = ""
	var isHeaderComplete bool = false

	for !self.isShutdown {
		curReadSize, readErr := self.conn.Read(recvBuffer[readSize:])
		if readErr != nil {
			clearRecvData(&recvBuffer, &recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
			logger.Error(readErr.Error())
			continue
		}

		if len(recvBuffer)+curReadSize >= PACKET_MAX_BUFFER {
			clearRecvData(&recvBuffer, &recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
			logger.Error("RecvBuffer overflow!")
			continue
		}

		readSize += curReadSize

		if !isHeaderComplete {
			// Can't Parse Header
			if readSize < PACKET_HEADERSIZE {
				continue
			}

			binaryReader := NewBinaryReader(recvBuffer)
			packetSize, _ = binaryReader.ReadInt32()

			isHeaderComplete = true
		}

		if isHeaderComplete {
			if packetSize < int32(readSize) {
				continue
			}
		}

		unmarshalErr := json.Unmarshal(recvBuffer, &recvPacket)
		if unmarshalErr != nil {
			clearRecvData(&recvBuffer, &recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
			logger.Error(unmarshalErr.Error())
			continue
		}

		self.recvChan <- recvPacket

		clearRecvData(&recvBuffer, &recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
	}
}

func clearRecvData(InBuffer *[]byte, InPacket *MessagePacket, InReadSize *int, InPacketSize *int32, InMessageData *string, InIsHeaderComplete *bool) {
	*InBuffer = (*InBuffer)[:0]
	*InPacket = MessagePacket{0, ""}
	*InReadSize = 0
	*InPacketSize = 0
	*InMessageData = ""
	*InIsHeaderComplete = false
}
