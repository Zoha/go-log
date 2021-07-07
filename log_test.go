package log

import (
	"strings"
	"testing"
)

type TestCase struct {
	log             string
	logF            string
	expected        string
	formatVariables []interface{}
	expectedF       string
}

type methodDetailsItem struct {
	method  func(a ...interface{})
	methodF func(format string, a ...interface{})
	level   uint8
	name    string
}

var testCases = []TestCase{
	{
		log:             "Hello",
		logF:            "Hello %v",
		expected:        "TestLogger: Hello\n",
		formatVariables: []interface{}{"Zoha"},
		expectedF:       "TestLogger: Hello Zoha",
	},
}

func TestLogger(t *testing.T) {
	log := Begin()

	// save all the methods in an array to call dinamically
	var methodsList = []methodDetailsItem{
		{
			method:  log.Alert,
			methodF: log.AlertF,
			level:   1,
			name:    "Alert",
		},
		{
			method:  log.Error,
			methodF: log.ErrorF,
			level:   1,
			name:    "Error",
		},
		{
			method:  log.Warn,
			methodF: log.WarnF,
			level:   2,
			name:    "Warn",
		},
		{
			method:  log.Highlight,
			methodF: log.HighlightF,
			level:   3,
			name:    "Highlight",
		},
		{
			method:  log.Inform,
			methodF: log.InformF,
			level:   4,
			name:    "Inform",
		},
		{
			method:  log.Log,
			methodF: log.LogF,
			level:   5,
			name:    "Log",
		},
		{
			method:  log.Trace,
			methodF: log.TraceF,
			level:   6,
			name:    "Trace",
		},
	}

	// for each test case test all the methods on logger
	for _, testCase := range testCases {
		for _, methodDetails := range methodsList {
			// declare variables and set destination of log
			var output strings.Builder
			var loggedString string
			log.logDestination = &output

			// log with lower level
			log.Level(methodDetails.level - 1)
			methodDetails.method(testCase.log)
			loggedString = output.String()
			if loggedString != "" {
				t.Error("logged string in lower level should be empty")
			}

			output.Reset()

			// logF with lower level
			log.Level(methodDetails.level - 1)
			methodDetails.methodF(testCase.logF, testCase.formatVariables...)
			loggedString = output.String()
			if loggedString != "" {
				t.Error("logged string in lower level should be empty")
			}

			output.Reset()

			// log with equal level
			log.Level(methodDetails.level)
			methodDetails.method(testCase.log)
			loggedString = output.String()
			if loggedString != testCase.expected {
				t.Errorf("for method %v expected %v but received %v", methodDetails.name, testCase.expected, loggedString)
			}

			output.Reset()

			// logF with equal level
			log.Level(methodDetails.level)
			methodDetails.methodF(testCase.logF, testCase.formatVariables...)
			loggedString = output.String()
			if loggedString != testCase.expectedF {
				t.Errorf("for format method %v expected %v but received %v", methodDetails.name, testCase.expectedF, loggedString)
			}

			output.Reset()

			// log with equal level
			log.Level(methodDetails.level + 1)
			methodDetails.method(testCase.log)
			loggedString = output.String()
			if loggedString != testCase.expected {
				t.Errorf("for method %v expected %v but received %v", methodDetails.name, testCase.expected, loggedString)
			}

			output.Reset()

			// logF with equal level
			log.Level(methodDetails.level + 1)
			methodDetails.methodF(testCase.logF, testCase.formatVariables...)
			loggedString = output.String()
			if loggedString != testCase.expectedF {
				t.Errorf("for format method %v expected %v but received %v", methodDetails.name, testCase.expectedF, loggedString)
			}
		}

	}

}
