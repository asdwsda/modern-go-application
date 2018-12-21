package greetingadapter

import (
	"fmt"
	"testing"

	"github.com/goph/logur"
	"github.com/goph/logur/testing"
	"github.com/sagikazarmark/modern-go-application/internal/greeting"
)

func TestLogger_Levels(t *testing.T) {
	tests := map[string]struct {
		logFunc func(logger *Logger, msg string, fields map[string]interface{})
	}{
		"trace": {
			logFunc: (*Logger).Trace,
		},
		"debug": {
			logFunc: (*Logger).Debug,
		},
		"info": {
			logFunc: (*Logger).Info,
		},
		"warn": {
			logFunc: (*Logger).Warn,
		},
		"error": {
			logFunc: (*Logger).Error,
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			testLogger := logur.NewTestLogger()
			logger := NewLogger(testLogger)

			test.logFunc(logger, fmt.Sprintf("message: %s", name), nil)

			level, _ := logur.ParseLevel(name)

			event := logur.LogEvent{
				Level: level,
				Line:  "message: " + name,
			}

			logtesting.AssertLogEvents(t, event, *(testLogger.LastEvent()))
		})
	}
}

func TestLogger_WithFields(t *testing.T) {
	testLogger := logur.NewTestLogger()

	var logger greeting.Logger = NewLogger(testLogger)

	fields := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	logger = logger.WithFields(fields)

	logger.Debug("message", nil)

	event := logur.LogEvent{
		Level:  logur.Debug,
		Line:   "message",
		Fields: fields,
	}

	logtesting.AssertLogEvents(t, event, *(testLogger.LastEvent()))
}
