package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Logger struct {
	logDestination io.Writer
}

func (l Logger) getLogDestination() io.Writer {
	// use l.logDestination or io.Discard by default for log destination
	logDestination := l.logDestination
	if logDestination == nil {
		logDestination = io.Discard
	}
	return logDestination
}

func (l Logger) Log(a ...interface{}) {
	logDestination := l.getLogDestination()
	fmt.Fprintln(logDestination, a...)
}

func (l Logger) LogF(format string, a ...interface{}) {
	logDestination := l.getLogDestination()
	fmt.Fprintf(logDestination, format, a...)
}

func (l *Logger) Begin() {
	l.logDestination = os.Stdout
	l.Log("BEGIN")
}

func (l *Logger) End() {
	l.Log("END")
	l.logDestination = ioutil.Discard
}
