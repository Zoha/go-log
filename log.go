package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	logDestination io.Writer
	prefix         string
	startTime      time.Time
	level          uint8
}

func getCallerFuncName() string {
	// get current go file name
	_, currentFileName, _, _ := runtime.Caller(0)

	var pc uintptr
	var callerFileName string
	skipIndex := 1
	for {
		skipIndex += 1
		pc, callerFileName, _, _ = runtime.Caller(skipIndex)
		// if is this file (logged from this package funcs) skip to another func
		if currentFileName != callerFileName || skipIndex > 5 {
			break
		}
	}

	caller := runtime.FuncForPC(pc)

	// convert something.else.funcName to funcName
	name := caller.Name()
	callersStack := strings.Split(name, ".")
	callerFuncName := callersStack[len(callersStack)-1]
	return callerFuncName
}

func getLogDestination(l Logger) io.Writer {
	// use l.logDestination or io.Discard by default for log destination
	logDestination := l.logDestination
	if logDestination == nil {
		logDestination = io.Discard
	}
	return logDestination
}

func log(l Logger, a ...interface{}) {
	logDestination := getLogDestination(l)

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

func logF(l Logger, format string, a ...interface{}) {
	logDestination := getLogDestination(l)
	funcName := getCallerFuncName()
	prefix := funcName + ": " + l.prefix
	fmt.Fprintf(logDestination, prefix+format, a...)
}

func Begin() Logger {
	logger := Logger{
		logDestination: os.Stdout,
		startTime:      time.Now(),
		level:          0,
	}
	logger.Trace("BEGIN")
	return logger
}

func (l *Logger) End() {
	// reset the prefix
	l.prefix = ""
	l.level = 6
	executionMicroseconds := time.Since(l.startTime).Microseconds()
	l.TraceF("END δt=%vµs\n", executionMicroseconds)
	l.logDestination = ioutil.Discard
}

func (l *Logger) Prefix(prefixes ...string) Logger {
	joinedStr := ""
	for _, str := range prefixes {
		joinedStr += str + ": "
	}
	l.prefix = joinedStr

	return *l
}

func (l *Logger) Level(level uint8) Logger {
	l.level = level
	return *l
}

func (l *Logger) Alert(a ...interface{}) {
	if l.level >= 1 {
		log(*l, a...)
	}
}

func (l *Logger) AlertF(format string, a ...interface{}) {
	if l.level >= 1 {
		logF(*l, format, a...)
	}
}

func (l *Logger) Error(a ...interface{}) {
	if l.level >= 1 {
		log(*l, a...)
	}
}

func (l *Logger) ErrorF(format string, a ...interface{}) {
	if l.level >= 1 {
		logF(*l, format, a...)
	}
}

func (l *Logger) Warn(a ...interface{}) {
	if l.level >= 2 {
		log(*l, a...)
	}
}

func (l *Logger) WarnF(format string, a ...interface{}) {
	if l.level >= 2 {
		logF(*l, format, a...)
	}
}

func (l *Logger) Highlight(a ...interface{}) {
	if l.level >= 3 {
		log(*l, a...)
	}
}

func (l *Logger) HighlightF(format string, a ...interface{}) {
	if l.level >= 3 {
		logF(*l, format, a...)
	}
}

func (l *Logger) Inform(a ...interface{}) {
	if l.level >= 4 {
		log(*l, a...)
	}
}

func (l *Logger) InformF(format string, a ...interface{}) {
	if l.level >= 4 {
		logF(*l, format, a...)
	}
}

func (l *Logger) Log(a ...interface{}) {
	if l.level >= 5 {
		log(*l, a...)
	}
}

func (l *Logger) LogF(format string, a ...interface{}) {
	if l.level >= 5 {
		logF(*l, format, a...)
	}
}

func (l *Logger) Trace(a ...interface{}) {
	if l.level >= 6 {
		log(*l, a...)
	}
}

func (l *Logger) TraceF(format string, a ...interface{}) {
	if l.level >= 6 {
		logF(*l, format, a...)
	}
}
