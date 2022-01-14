package main

import (
	"chat/corelib/logger"
	"chat/execute"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	execute.StartClient()

	logger.Text("===== Client Start =====")

	execute.RunClient()

	logger.Text("===== Client Shutdown =====")
}
