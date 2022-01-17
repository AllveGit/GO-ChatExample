package netclient

import (
	"bufio"
	"chat/corelib/logger"
	"chat/corelib/network"
	. "chat/corelib/network/session"
	. "chat/types"
	"os"
	"os/signal"
	"strings"
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
			KillSignalChan <- syscall.SIGTERM
			return
		}

		logger.Text("connect success to %s", conn.RemoteAddr().String())

		session, sessionMakeErr := NewPacketSession(0, conn)
		if sessionMakeErr != nil {
			logger.Error(sessionMakeErr.Error())
		}
		defer func() {
			logger.Error("Session End!")
			session.Shutdown()
			KillSignalChan <- syscall.SIGTERM
		}()

		for {
			in := bufio.NewReader(os.Stdin)
			message, inputErr := in.ReadString('\n')
			if inputErr != nil {
				logger.Error(inputErr.Error())
				return
			}

			if strings.Compare(message, "quit\r\n") == 0 {
				break
			}

			messagePacket := MessagePacket{int32(len(message)), message}
			session.Send(messagePacket)
		}
	}()

	<-KillSignalChan
	logger.Warning("KillSignal Received")
}
