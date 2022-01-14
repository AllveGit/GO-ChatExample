package logger

import (
	"fmt"
	"io"
	"time"
)

var (
	_DebugEnabled   = true
	_ImmediateMode  = true
	_UseColor       = true
	_LogLevel       = 0
	_LogWriter      io.Writer
	_LogInfoChannel chan LogInfo
)

type RealLogWriter struct {
}

func (self *RealLogWriter) Write(InBuf []byte) (int, error) {
	fmt.Printf("%s\n", InBuf)
	return len(InBuf), nil
}

func Initialize() {
	_LogInfoChannel = make(chan LogInfo, 4096)
	_LogWriter = new(RealLogWriter)
	_UseColor = true

	go func() {
		for {
			l := <-_LogInfoChannel
			PrintLogInfo(&l)
		}
	}()
}

func AddLogInfo(l LogInfo) {
	if _ImmediateMode {
		PrintLogInfo(&l)
	} else {
		_LogInfoChannel <- l
	}
}

func PrintLogInfo(l *LogInfo) {
	if l.Level >= _LogLevel {
		now := time.Now()
		logMsg := fmt.Sprintf("[%d/%d %d:%d:%d] %s%s%s", now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), l.logType.ToString(), l.SignifyText, l.NormalText)

		if _UseColor && l.logType.HasColor() {
			l.logType.SetColor()
			fmt.Fprint(_LogWriter, logMsg)
			l.logType.ResetColor()
		} else {
			fmt.Fprint(_LogWriter, logMsg)
		}
	}
}
