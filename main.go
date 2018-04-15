package dlog

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	datadogFormatter := &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "timestamp",
			log.FieldKeyLevel: "level",
			log.FieldKeyMsg:   "message",
		},
	}
	log.SetFormatter(datadogFormatter)
	log.SetOutput(os.Stdout)
}

const (
	RequestIO = "X-Request-ID"
	Service   = "service"
)

func NewLogger(service string) *log.Entry {
	return log.WithFields(log.Fields{
		Service: &service,
	})
}

func NewRequestLogger(requestId string, service string) *log.Entry {
	return log.WithFields(log.Fields{
		RequestIO: &requestId,
		Service:   &service,
	})
}
