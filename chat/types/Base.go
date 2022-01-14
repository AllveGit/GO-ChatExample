package types

type Port uint16

// Simple message packet.
type MessagePacket struct {
	PacketLen int32
	Message   string
}
