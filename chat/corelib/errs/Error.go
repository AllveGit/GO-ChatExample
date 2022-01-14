package errs

import (
	"fmt"
	"path"
	"runtime"
)

func New(InFormat string, InArgs ...interface{}) error {
	var fileName, funcName string

	pc, filePath, line, ok := runtime.Caller(1)
	if ok {
		fileName = path.Base(filePath)
		funcName = runtime.FuncForPC(pc).Name()
	}

	msg := fmt.Sprintf(InFormat, InArgs...)
	err := fmt.Errorf("%s \n\tat %s() %s:%d", msg, funcName, fileName, line)

	return err
}
