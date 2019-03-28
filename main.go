package dlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func init() {
	datadogFormatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	}
	logrus.SetFormatter(datadogFormatter)
	logrus.SetOutput(os.Stdout)
}

const (
	// CorrelationID used to correlate logs
	CorrelationID = "correlationid"

	// AppName visible as a service
	AppName = "appname"
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
}

// NewStandardLogger creates standard logger
func NewStandardLogger(loggerParam *LoggerParam) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		CorrelationID: &loggerParam.CorrelationID,
		AppName:       &loggerParam.AppName,
	})
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
	splitted := strings.Split(arn, ":")
	appName := splitted[len(splitted)-1]
	return appName, nil
}
