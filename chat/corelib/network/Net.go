package network

import "net"

var (
	_Network INetwork
)

func Initialize() {
	_Network = new(TCPNetwork)
}

func Listen(InAddr string) (net.Listener, error) {
	return _Network.Listen(InAddr)
}

func Connect(InAddr string) (net.Conn, error) {
	return _Network.Connect(InAddr)
}
