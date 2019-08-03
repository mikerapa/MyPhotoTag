package cli

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

const dateTimeFormat = "2 Jan 2006 15:04:05"

//Logger
var ConsoleLogger logrus.Logger

type myFormatter struct {
	logrus.TextFormatter
}

func (f *myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//formattedLogEntry := fmt.Sprintln(entry.Time.Format(dateTimeFormat), strings.ToUpper(entry.Level.String()), "-", entry.Message)
	formattedLogEntry := fmt.Sprintln(strings.ToUpper(entry.Level.String()), "-", entry.Message)
	byteBuffer := &bytes.Buffer{}
	byteBuffer.WriteString(formattedLogEntry)
	return byteBuffer.Bytes(), nil
}

func getFormatter() logrus.Formatter {
	formatter := new(myFormatter)
	formatter.ForceColors = true
	return formatter
}

// Set up the logger instance for use
func InitLogger(logLevelString string) {
	ConsoleLogger = GetLogger(GetLogLevelFromString(logLevelString))
}

func GetLogger(logLevel logrus.Level) logrus.Logger {
	logger := logrus.Logger{
		Out:       os.Stderr,
		Level:     logLevel,
		Formatter: getFormatter(),
	}

	return logger
}

// (error, warning, info, debug, trace)
func GetLogLevelFromString(logLevelString string) logrus.Level {
	switch strings.ToLower(logLevelString) {
	case "info":
		return logrus.InfoLevel
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.FatalLevel
	}
}
