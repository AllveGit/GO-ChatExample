package network

import (
	"net"
	"time"
)

type INetwork interface {
	Listen(InAddr string) (net.Listener, error)
	Connect(InAddr string) (net.Conn, error)
}

type TCPNetwork struct {
}

func (self *TCPNetwork) Listen(InAddr string) (net.Listener, error) {
	return net.Listen("tcp", InAddr)
}

func (self *TCPNetwork) Connect(InAddr string) (net.Conn, error) {
	return net.DialTimeout("tcp", InAddr, time.Second*5)
}
