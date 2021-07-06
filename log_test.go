package log

import (
	"strings"
	"testing"
)

func TestLogger_Log(t *testing.T) {
	var output strings.Builder

	msgToLog := "my message to log"
	expectedLog := "TestLogger_Log: " + msgToLog + "\n"

	log := Logger{}

	log.logDestination = &output
	log.Log(msgToLog)

	if output.String() != expectedLog {
		t.Errorf("expected %s but %s received", expectedLog, output.String())
	}

}

func TestLogger_LogF(t *testing.T) {
	var output strings.Builder

	name := "Zoha"
	formatToLog := "Hello %v"
	expectedLog := "TestLogger_LogF: " + "Hello " + name

	log := Logger{}

	log.logDestination = &output
	log.LogF(formatToLog, name)

	outputString := output.String()

	if outputString != expectedLog {
		t.Errorf("expected %s but %s received", expectedLog, outputString)
	}

}
