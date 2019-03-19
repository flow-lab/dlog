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
	RequestId = "X-Request-ID"
	Service   = "service"
)

func NewLogger(service string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		Service: &service,
	})
}

func NewRequestLogger(requestId string, service string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		RequestId: &requestId,
		Service:   &service,
	})
}
