package dlog

import (
	"github.com/sirupsen/logrus"
	"os"
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
	RequestIO = "X-Request-ID"
	AppName   = "appname"
)

func NewLogger(service string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		AppName: &service,
	})
}

func NewRequestLogger(requestId string, service string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		RequestIO: &requestId,
		AppName:   &service,
	})
}
