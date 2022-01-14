package logger

import (
	"fmt"
	"path"
	"runtime"
)

func ImmediateMode(InIsEnable bool) {
	_ImmediateMode = InIsEnable
}

func EnableDebug(InIsEnable bool) {
	_DebugEnabled = InIsEnable
}

func SetLogLevel(InLevel int) {
	_LogLevel = InLevel
}

func StackBuffer() []byte {
	var buffer []byte
	bufferSize := 2048

	for {
		buffer = make([]byte, bufferSize)
		n := runtime.Stack(buffer, false)
		if n <= bufferSize {
			break
		}

		bufferSize *= 2
	}

	return buffer
}

func Text(InFormat string, InArgs ...interface{}) {
	l := LogInfo{
		LOG_TYPE_TEXT,
		"",
		fmt.Sprintf(InFormat, InArgs...),
		0,
	}

	AddLogInfo(l)
}

func Debug(InFormat string, InArgs ...interface{}) {
	if !_DebugEnabled {
		return
	}

	// color msg
	var signifyMsg string

	pc, filePath, line, ok := runtime.Caller(1)
	if ok {
		fileName := path.Base(filePath)
		funcName := runtime.FuncForPC(pc).Name()
		signifyMsg = fmt.Sprintf("%s() in %s:%d", funcName, fileName, line)
	}

	l := LogInfo{
		LOG_TYPE_DEBUG,
		signifyMsg,
		fmt.Sprintf(InFormat, InArgs...),
		0,
	}

	AddLogInfo(l)
}

func DebugLogLevel(InLevel int, InFormat string, InArgs ...interface{}) {
	// color msg
	var signifyMsg, normalMsg string

	normalMsg = fmt.Sprintf(InFormat, InArgs...)

	pc, filePath, line, ok := runtime.Caller(1)
	if ok {
		fileName := path.Base(filePath)
		funcName := runtime.FuncForPC(pc).Name()
		signifyMsg = fmt.Sprintf("%s() in %s:%d", funcName, fileName, line)
	}

	l := LogInfo{
		LOG_TYPE_ERROR,
		signifyMsg,
		normalMsg,
		InLevel,
	}

	AddLogInfo(l)
}

func Warning(InFormat string, InArgs ...interface{}) {
	l := LogInfo{
		LOG_TYPE_WARNING,
		"",
		fmt.Sprintf(InFormat, InArgs...),
		0,
	}

	AddLogInfo(l)
}

func Error(InFormat string, InArgs ...interface{}) {
	// color msg
	var signifyMsg string

	pc, filePath, line, ok := runtime.Caller(1)
	if ok {
		fileName := path.Base(filePath)
		funcName := runtime.FuncForPC(pc).Name()
		signifyMsg = fmt.Sprintf("%s() in %s:%d", funcName, fileName, line)
	}

	l := LogInfo{
		LOG_TYPE_ERROR,
		signifyMsg,
		fmt.Sprintf(InFormat, InArgs...),
		0,
	}

	AddLogInfo(l)
}

func Critical(InFormat string, InArgs ...interface{}) {
	// callstack
	buffer := StackBuffer()

	callStack := fmt.Sprint(
		"\n---------- callstack begin ----------\n",
		string(buffer),
		"---------- callstack end ----------\n",
	)

	l := LogInfo{
		LOG_TYPE_CRITICAL,
		fmt.Sprintf(InFormat, InArgs...),
		callStack,
		0,
	}

	AddLogInfo(l)
}
