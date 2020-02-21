package dlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type hook struct{}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *hook) Fire(e *logrus.Entry) error {
	for k, v := range e.Data {
		if s, ok := v.(*string); ok {
			if *s == "" {
				delete(e.Data, k)
				continue
			}
		}
	}
	return nil
}

func init() {
	f := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	}
	logrus.SetFormatter(f)
	logrus.SetOutput(os.Stdout)

	logrus.AddHook(&hook{})
}

const (
	// CorrelationID used to correlate logs
	CorrelationID = "correlationid"

	// AppName visible as a service
	AppName = "appname"

	// Parent segment trace id
	Parent = "parent"

	// Trace current trace id
	Trace = "trace"

	// Span id
	Span = "span"

	// Version of the application
	Version = "version"

	// Commit Short SHA
	Commit = "commit"

	// Build info
	Build = "build"
)

// NewLogger creates standard logger
func NewLogger(appName string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		AppName: &appName,
	})
}

// LoggerParam parameters for creating a logger
type LoggerParam struct {
	CorrelationID string
	AppName       string
	Parent        string
	Trace         string
	Span          string
	Version       string
	Commit        string
	Build         string
}

// NewStandardLogger creates standard logger
func NewStandardLogger(loggerParam *LoggerParam) *logrus.Entry {
	fields := logrus.WithFields(logrus.Fields{
		CorrelationID: &loggerParam.CorrelationID,
		AppName:       &loggerParam.AppName,
		Parent:        &loggerParam.Parent,
		Trace:         &loggerParam.Trace,
		Span:          &loggerParam.Span,
		Version:       &loggerParam.Version,
		Commit:        &loggerParam.Commit,
		Build:         &loggerParam.Build,
	})
	return fields
}

// NewRequestLogger creates standard logger with correlationId and appName
// Deprecated: Use strings.HasPrefix instead.
func NewRequestLogger(correlationID string, service string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		CorrelationID: &correlationID,
		AppName:       &service,
	})
}

// GetAppNameFromARN gets for example lambda name from ARN
func GetAppNameFromARN(arn string) (string, error) {
	if arn == "" {
		return "", fmt.Errorf("arn cannot be blank")
	}
	s := strings.Split(arn, ":")
	return s[len(s)-1], nil
}
