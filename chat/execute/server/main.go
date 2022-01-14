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

	execute.StartServer()

	logger.Text("===== Server Open =====")

	execute.RunServer()

	logger.Text("===== Server Close =====")
}
