package main

import (
	"chat/corelib/logger"
	"chat/corelib/network"
	"chat/corelib/network/netclient"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	Start()

	logger.Text("===== Client Start =====")

	Run()

	logger.Text("===== Client Shutdown =====")
}

func Start() {
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

func Run() {
	netclient.Run()
}
