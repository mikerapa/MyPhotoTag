package cli

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetLogger(t *testing.T) {
	tests := []struct {
		name          string
		inputLogLevel logrus.Level
	}{
		{"Debug", logrus.DebugLevel},
		{"Trace", logrus.TraceLevel},
		{"Fatal", logrus.FatalLevel},
		{"Info", logrus.InfoLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLogger(tt.inputLogLevel)
			if got.Level != tt.inputLogLevel {
				t.Errorf("GetLogger() did not create a correct logger. %v", got)
			}
		})
	}
}

func TestGetLogLevelFromString(t *testing.T) {
	tests := []struct {
		name           string
		logLevelString string
		want           logrus.Level
	}{
		{"Info", "Info", logrus.InfoLevel},
		{"info", "info", logrus.InfoLevel},
		{"Empty", "", logrus.FatalLevel},
		{"fatal", "fatal", logrus.FatalLevel},
		{"DEBUG", "DEBUG", logrus.DebugLevel},
		{"Error", "Error", logrus.ErrorLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLogLevelFromString(tt.logLevelString); got != tt.want {
				t.Errorf("GetLogLevelFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name           string
		logLevelString string
		want           logrus.Level
	}{
		{"Info", "Info", logrus.InfoLevel},
		{"Error", "error", logrus.ErrorLevel},
		{"Debug", "DEBUG", logrus.DebugLevel},
		{"Trace", "trace", logrus.TraceLevel},
		{"FAtal", "FATAL", logrus.FatalLevel},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogger(tt.logLevelString)
			if ConsoleLogger.Level != tt.want {
				t.Errorf("InitLogger faile to set up the logger. %v", tt.want)
			}
		})
	}
}
