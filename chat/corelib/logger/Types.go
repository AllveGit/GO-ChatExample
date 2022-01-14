package logger

import (
	ct "github.com/daviddengcn/go-colortext"
)

type logType int32

type LogInfo struct {
	logType     logType
	SignifyText string
	NormalText  string
	Level       int
}

const (
	LOG_TYPE_TEXT logType = iota
	LOG_TYPE_DEBUG
	LOG_TYPE_WARNING
	LOG_TYPE_ERROR
	LOG_TYPE_CRITICAL
)

func (l logType) ToString() string {
	switch l {
	case LOG_TYPE_CRITICAL:
		return "[CRITICAL] "
	case LOG_TYPE_ERROR:
		return "[ERROR] "
	case LOG_TYPE_WARNING:
		return "[WARNING] "
	case LOG_TYPE_DEBUG:
		return "[DEBUG] "
	default:
		return ""
	}
}

func (l logType) SetColor() {
	switch l {
	case LOG_TYPE_CRITICAL:
		ct.ChangeColor(ct.Red, true, ct.Yellow, false)
	case LOG_TYPE_ERROR:
		ct.ChangeColor(ct.Yellow, true, ct.None, false)
	case LOG_TYPE_WARNING:
		ct.ChangeColor(ct.Cyan, true, ct.None, false)
	case LOG_TYPE_DEBUG:
		ct.ChangeColor(ct.Green, true, ct.Magenta, false)
	}
}

func (l logType) ResetColor() {
	ct.ResetColor()
}

func (l logType) HasColor() bool {
	switch l {
	case LOG_TYPE_TEXT:
		return false
	default:
		return true
	}
}
