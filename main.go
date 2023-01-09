package dlog

import (
	"github.com/sirupsen/logrus"
	"os"
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

	// Func is the name of the function
	Func = "func"

	// File is the file name with line number
	File = "file"
)

// Config parameters for creating a logger
type Config struct {
	CorrelationID string
	AppName       string
	Parent        string
	Trace         string
	Span          string
	Version       string
	Commit        string
	Build         string
	Level         string
	ReportCaller  bool
}

// NewLogger creates logger based on the config
func NewLogger(c *Config) *logrus.Entry {
	if c == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}
	logrus.SetReportCaller(c.ReportCaller)

	if c.Level != "" {
		parseLevel, err := logrus.ParseLevel(c.Level)
		if err == nil {
			logrus.SetLevel(parseLevel)
		}
	}

	fields := logrus.WithFields(logrus.Fields{
		CorrelationID: &c.CorrelationID,
		AppName:       &c.AppName,
		Parent:        &c.Parent,
		Trace:         &c.Trace,
		Span:          &c.Span,
		Version:       &c.Version,
		Commit:        &c.Commit,
		Build:         &c.Build,
	})
	return fields
}
