## dlog - Datadog Go logger ![Go](https://github.com/flow-lab/dlog/workflows/Go/badge.svg) [![codecov](https://codecov.io/gh/flow-lab/dlog/branch/master/graph/badge.svg)](https://codecov.io/gh/flow-lab/dlog) [![Go Report Card](https://goreportcard.com/badge/github.com/flow-lab/dlog)](https://goreportcard.com/report/github.com/flow-lab/dlog)

Go logger which logs messages in [Datadog](https://docs.datadoghq.com/logs/)
json format. Build with https://github.com/sirupsen/logrus

## Context logger

```go
import (
  ...
  log "github.com/sirupsen/logrus"
  "github.com/flow-lab/dlog"
)

...

logger := dlog.NewStandardLogger(&LoggerParam{
		AppName:       "MyService",
		Trace:         "1-5d0a8b05-4d6952b21901d9396e578955",
		Parent:        "0c1db8f76a4f6073",
})

logger := dlog.NewLogger(&dlog.Config{
    AppName:      "myservice",
    Level:        "debug",
    Version:      "0.1.0",
    Commit:       "1234567",
    Build:        "2020-01-01T00:00:00Z",
    ReportCaller: true,
})

logger.Info("Hello world")
{"appname":"myservice","build":"2020-01-01T00:00:00Z","commit":"1234567","file":"/Users/test/dlog/main_test.go:82","func":"github.com/flow-lab/dlog.TestContextLogger.func2","level":"info","message":"Hello World","timestamp":"2023-01-09T16:17:36+01:00","version":"0.1.0"}
```

License
-------
[![License: MIT](https://img.shields.io/badge/License-mit-brightgreen.svg)](https://opensource.org/licenses/MIT)
