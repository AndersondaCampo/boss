package util

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ErrorList struct {
	errList []error
}

func (e *ErrorList) AppendErr(err error) {
	e.errList = append(e.errList, err)
}

func (e *ErrorList) Error() string {
	var result string
	for _, err := range e.errList {
		result += fmt.Sprintf("%s\n", err.Error())
	}
	return result
}

func (e *ErrorList) HasErr() bool {
	return len(e.errList) > 0
}

func (e *ErrorList) Result() error {
	if e.HasErr() {
		return e
	} else {
		return nil
	}
}

func fatal(err error) {
	logrus.Panic(err)
}

func err(err error) {
	logrus.Error(err)
}

var FatalErrHandler = fatal
var ErrorErrHandler = err

const DefaultErrorExitCode = 1

func CheckErr(err error) {
	_ = checkError(err, FatalErrHandler)
}

func CheckErrSoft(err error) bool {
	return checkError(err, ErrorErrHandler)
}

func checkError(err error, handleErr func(error)) bool {
	if err == nil {
		return false
	}

	handleErr(err)
	return true
}

func UsageErrorf(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nSee '%s -h' for help and examples", msg, cmd.CommandPath())
}
