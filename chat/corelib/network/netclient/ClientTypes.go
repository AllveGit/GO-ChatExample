package netclient

import (
	"chat/corelib/logger"
	"chat/corelib/network"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	ServerAddress = "127.0.0.1:9999"
)

var KillSignalChan chan os.Signal

type IClient interface {
	Run()
}

type ChatClient struct {
}

func (self *ChatClient) Run() {
	KillSignalChan = make(chan os.Signal)
	signal.Notify(KillSignalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		conn, connErr := network.Connect("127.0.0.1:9999")
		if connErr != nil {
			logger.Error(connErr.Error())
		}

		go func(InConn net.Conn) {
			logger.Text("connect suscess to %s", conn.RemoteAddr().String())
			conn.Close()
			KillSignalChan <- syscall.SIGTERM
		}(conn)
	}()

	<-KillSignalChan
	logger.Warning("KillSignal Received")
}
