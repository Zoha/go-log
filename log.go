package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	logDestination io.Writer
	prefix         string
}

func getCallerFuncName() string {
	// get caller func
	pc, callerFileName, _, _ := runtime.Caller(2)
	// get current go file name
	_, currentFileName, _, _ := runtime.Caller(0)

	// if is this file (logged from this package funcs) skip to another func
	if currentFileName == callerFileName {
		pc, _, _, _ = runtime.Caller(3)
	}

	caller := runtime.FuncForPC(pc)

	// convert something.else.funcName to funcName
	name := caller.Name()
	callersStack := strings.Split(name, ".")
	callerFuncName := callersStack[len(callersStack)-1]
	return callerFuncName
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

	funcName := getCallerFuncName()
	prefix := funcName + ":"
	if l.prefix != "" {
		// we add  an extra space for extra args space (alignment)
		prefix += " " + l.prefix
	}

	args := a
	// prepend prefixes (func name and other prefixes) to args
	args = append(a, 0)
	copy(args[1:], args)
	args[0] = prefix

	fmt.Fprintln(logDestination, args...)
}

func (l Logger) LogF(format string, a ...interface{}) {
	logDestination := l.getLogDestination()
	funcName := getCallerFuncName()
	prefix := funcName + ": " + l.prefix
	fmt.Fprintf(logDestination, prefix+format, a...)
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
