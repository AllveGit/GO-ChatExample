package session

import (
	. "chat/config"
	"chat/corelib/logger"
	. "chat/corelib/network"
	. "chat/types"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type PacketSession struct {
	id             int
	conn           net.Conn
	isShutdown     bool
	sendChan       chan MessagePacket
	recvChan       chan MessagePacket
	SessionManager ISessionManager
}

func NewPacketSession(InID int, InConn net.Conn) (*PacketSession, error) {
	session := &PacketSession{
		id:             InID,
		conn:           InConn,
		isShutdown:     false,
		sendChan:       make(chan MessagePacket, 10),
		recvChan:       make(chan MessagePacket, 10),
		SessionManager: nil,
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
		} else {
			logger.Text("session[%d] shutdowned.", self.id)
		}
	}
}

func (self *PacketSession) startPacketEvent() {
	for !self.isShutdown {
		select {
		case sendPkt := <-self.sendChan:
			sendBytes, marshalErr := json.Marshal(sendPkt)
			if marshalErr != nil {
				logger.Error(marshalErr.Error())
				continue
			}

			self.conn.Write(sendBytes)
		case recvPkt := <-self.recvChan:
			fmt.Printf(recvPkt.Message)

			if self.SessionManager != nil {
				self.SessionManager.BroadCast(self.id, recvPkt)
			}
		}
	}
}

func (self *PacketSession) startReceive() {

	defer self.Shutdown()

	recvBuffer := make([]byte, PACKET_MAX_BUFFER)
	recvPacket := MessagePacket{0, ""}

	var readSize int = 0
	var packetSize int32 = 0
	var messageData string = ""
	var isHeaderComplete bool = false

	for !self.isShutdown {
		curReadSize, readErr := self.conn.Read(recvBuffer[readSize:])
		if readErr != nil {
			if readErr == io.EOF {
				logger.Text("session[%d] disconnected.", self.id)
			} else {
				logger.Error(readErr.Error())
			}

			self.isShutdown = true
			clearRecvData(&recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
			continue
		}

		if curReadSize <= 0 {
			continue
		}

		if readSize+curReadSize >= PACKET_MAX_BUFFER {
			clearRecvData(&recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
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

		unmarshalErr := json.Unmarshal(recvBuffer[:readSize], &recvPacket)
		if unmarshalErr != nil {
			clearRecvData(&recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
			logger.Error(unmarshalErr.Error())
			continue
		}

		self.recvChan <- recvPacket
		clearRecvData(&recvPacket, &readSize, &packetSize, &messageData, &isHeaderComplete)
	}
}

func clearRecvData(InPacket *MessagePacket, InReadSize *int, InPacketSize *int32, InMessageData *string, InIsHeaderComplete *bool) {
	*InPacket = MessagePacket{0, ""}
	*InReadSize = 0
	*InPacketSize = 0
	*InMessageData = ""
	*InIsHeaderComplete = false
}
