package execute

import (
	"chat/corelib/logger"
	"chat/corelib/network"
	"chat/corelib/network/netclient"
	"chat/corelib/network/netserver"
)

func StartClient() {
	logger.EnableDebug(true)
	logger.ImmediateMode(true)
	logger.Initialize()

	logger.Text("===== Client Initialize Start =====")

	logger.Text("Logger Initialize Finish")

	// Network Initialize
	network.Initialize()

	logger.Text("Network Initialize Finish")

	// Client Initialize
	netclient.Initialize()

	logger.Text("NetClient Initialize Finish")
}

func StartServer() {
	// Logger Initialize
	logger.EnableDebug(true)
	logger.ImmediateMode(true)
	logger.Initialize()

	logger.Text("===== Server Initialize Start =====")

	logger.Text("Logger Initialize Finish")

	// Network Initialize
	network.Initialize()

	logger.Text("Network Initialize Finish")

	// Server Initialize
	netserver.Initialize()

	logger.Text("NetServer Initialize Finish")
}

func RunClient() {
	netclient.Run()
}

func RunServer() {
	netserver.Run()
}
