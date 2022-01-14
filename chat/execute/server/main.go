package main

import (
	"chat/corelib/logger"
	"chat/corelib/network"
	"chat/corelib/network/netserver"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	Start()

	logger.Text("===== Server Open =====")

	Run()

	logger.Text("===== Server Close =====")
}

func Start() {
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

func Run() {
	netserver.Run()
}
