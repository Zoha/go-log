package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Logger struct {
	logDestination io.Writer
	prefix         string
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

	args := append(a, 0)
	copy(args[1:], args)
	args[0] = l.prefix

	fmt.Fprintln(logDestination, args...)
}

func (l Logger) LogF(format string, a ...interface{}) {
	logDestination := l.getLogDestination()
	fmt.Fprintf(logDestination, l.prefix+format, a...)
}

func (l *Logger) Begin() {
	l.logDestination = os.Stdout
	l.Log("BEGIN")
}

func (l *Logger) End() {
	// reset the prefix
	l.prefix = ""
	l.Log("END")
	l.logDestination = ioutil.Discard
}

func (l *Logger) Prefix(prefixes ...string) {
	joinedStr := ""
	for _, str := range prefixes {
		joinedStr += str + ": "
	}
	l.prefix = joinedStr
}
