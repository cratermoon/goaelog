package elog

import (
	"fmt"

	"github.com/goadesign/goa"
)

type adapter struct {
	NulLogger
}

type NulLogger struct {
	count int
}

// New returns a goa log adapter backed by a nul logger.
func New(logger NulLogger) goa.LogAdapter {
	return &adapter{NulLogger: logger}
}

func NewNulLogger() NulLogger {
	return NulLogger{0}
}

func (a *adapter) Info(msg string, keyvals ...interface{}) {
	a.NulLogger.count++
	fmt.Printf("%d: %s %+v\n", a.NulLogger.count, msg, keyvals)
}
func (a *adapter) Error(msg string, keyvals ...interface{}) {
	a.NulLogger.count++
	fmt.Printf("%d: %s %+v\n", a.NulLogger.count, msg, keyvals)
}
func (a *adapter) New(keyvals ...interface{}) goa.LogAdapter {
	return &adapter{NulLogger: NulLogger{0}}
}
