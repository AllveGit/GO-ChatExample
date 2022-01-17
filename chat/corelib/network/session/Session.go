package session

import . "chat/types"

type ISession interface {
	Send(InPacket MessagePacket) error
	Shutdown()
}
