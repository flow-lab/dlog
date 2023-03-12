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
	CorrelationID string // correlation id
	AppName       string // name of the application, e.g. diatom
	Parent        string // parent id
	Trace         string // trace id
	Span          string // span id
	Version       string // version of the application, e.g. 0.1.0
	Commit        string // short SHA
	Build         string // build number, e.g. 123
	Level         string // debug, info, warn, error, fatal, panic
	ReportCaller  bool   // default is false
	Formatter     string // json or text, default is json
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

	if c.Formatter == "text" {
		// this should be used for local development
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors:          false,
			DisableLevelTruncation: true,
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05.000Z",
			PadLevelText:           true,
			ForceColors:            true,
		})
		logrus.SetOutput(os.Stdout)
		return logrus.NewEntry(logrus.StandardLogger())
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
