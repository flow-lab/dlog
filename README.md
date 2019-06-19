## dlog - Datadog Go logger [![Build Status](https://travis-ci.org/flow-lab/dlog.svg?branch=master)](https://travis-ci.org/flow-lab/dlog) [![codecov](https://codecov.io/gh/flow-lab/dlog/branch/master/graph/badge.svg)](https://codecov.io/gh/flow-lab/dlog)

Go logger which logs messages in [Datadog](https://docs.datadoghq.com/logs/)
json format. Build with https://github.com/sirupsen/logrus

## Logger with default application name

```go
import (
  ...
  log "github.com/sirupsen/logrus"
  "github.com/flow-lab/dlog"
)

...

logger := NewLogger("MyService")

logger.Info("Hello world")
{"level":"info","message":"Hello World","appname":"MyService","timestamp":"2018-04-15T21:06:00+02:00"}
```

## Context logger

```go
import (
  ...
  log "github.com/sirupsen/logrus"
  "github.com/flow-lab/dlog"
)

...

correlationID := "1-581cf771-a006649127e371903a2de979"
logger := NewStandardLogger(&LoggerParam{
		AppName:       "MyService",
		Trace:         "1-5d0a8b05-4d6952b21901d9396e578955",
		Parent:        "0c1db8f76a4f6073",
})

logger.Info("Hello world")
{"correlationid":"1-581cf771-a006649127e371903a2de979","level":"info","message":"Hello World","appname":"MyService","timestamp":"2018-04-15T21:05:19+02:00"}
```
