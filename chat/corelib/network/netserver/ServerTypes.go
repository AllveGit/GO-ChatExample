package netserver

import (
	"chat/corelib/logger"
	"chat/corelib/network"
	. "chat/corelib/network/session"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	ServerAddress = "127.0.0.1:9999"
	MAX_USER_NUM  = 64
)

var KillSignalChan chan os.Signal

type IServer interface {
	Run()
}

type ChatServer struct {
}

func (self *ChatServer) Run() {
	KillSignalChan = make(chan os.Signal)
	signal.Notify(KillSignalChan, syscall.SIGINT, syscall.SIGTERM)

	// Listen
	logger.Text("Listen Start!")

	listener, listenErr := network.Listen(ServerAddress)
	if listenErr != nil {
		logger.Error(listenErr.Error())
		return
	}

	// Accept
	logger.Text("Accept Start!")

	go func() {
		sessionManager := NewSessionManager()

		for {
			conn, acceptErr := listener.Accept()
			if acceptErr != nil {
				logger.Error(acceptErr.Error())
				break
			}

			go func(InConn net.Conn) {
				logger.Text("accepted new connect from %s", conn.RemoteAddr().String())

				_, newSessionErr := sessionManager.GenerateSession(InConn)
				if newSessionErr != nil {
					logger.Error(newSessionErr.Error())
					return
				}
			}(conn)
		}
	}()

	<-KillSignalChan
	logger.Warning("Server KillSignal Received")

	listener.Close()
}
