## dlog - Go logger ![Go](https://github.com/flow-lab/dlog/workflows/Go/badge.svg) [![codecov](https://codecov.io/gh/flow-lab/dlog/branch/master/graph/badge.svg)](https://codecov.io/gh/flow-lab/dlog) [![Go Report Card](https://goreportcard.com/badge/github.com/flow-lab/dlog)](https://goreportcard.com/report/github.com/flow-lab/dlog)

The _dlog_ package provides a logging framework for Go that formats logs in JSON format, optimised for use with Google
Kubernetes, AWS EKS etc. Logs can be easily pushed to Datadog, Google Stackdriver and AWS CloudWatch.

This package provides a simple and efficient way to format logs in a structured way that can be easily parsed by log
aggregators and monitoring tools, while still being human-readable. The JSON format also allows for easy searching and
filtering of logs, making it ideal for managing large-scale deployments.

In addition to the standard logging features, _dlog_ provides additional features like tagging logs with metadata and
filtering logs based on log levels. The package is highly customisable and easily integrated into any Go project.

Build on top of [Logrus](https://github.com/sirupsen/logrus). Log format matches json
format of [Datadog](https://docs.datadoghq.com/logs/).

## Installation

```shell
go get github.com/flow-lab/dlog
```

## Usage

### Basic Usage with JSON formatter

This format is used by default. It is optimised for use with Datadog, Google Stackdriver and AWS CloudWatch.

```go
import (
...
  log "github.com/sirupsen/logrus"
  "github.com/flow-lab/dlog"
)

...

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

logger := logger.WithField("component", "myprocessor")
logger.Info("Hello world")
{"appname":"myservice","component":"myprocessor","build":"2020-01-01T00:00:00Z","commit":"1234567","file":"/Users/test/dlog/main_test.go:82","func":"github.com/flow-lab/dlog.TestContextLogger.func2","level":"info","message":"Hello World","timestamp":"2023-01-09T16:17:36+01:00","version":"0.1.0"}
```


## Basic Usage with Text formatter

This format is used for local development.

```go
import (
...
  log "github.com/sirupsen/logrus"
  "github.com/flow-lab/dlog"
)

...

logger := dlog.NewLogger(&dlog.Config{
  AppName:      "myservice",
  Level:        "debug",
  Version:      "0.1.0",
  Commit:       "1234567",
  Build:        "2020-01-01T00:00:00Z",
  ReportCaller: false,
  Formatter:    "text", // for local development
})

logger.Info("Hello World")
logger.Debug("Hello World")
logger.Warn("Hello World")
logger.Error("Hello World")

INFO   [2023-03-12 23:04:11.377Z] Hello World
DEBUG  [2023-03-12 23:04:11.378Z] Hello World
WARNING[2023-03-12 23:04:11.378Z] Hello World
ERROR  [2023-03-12 23:04:11.378Z] Hello World
```

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue on GitHub. If you want to
contribute code, please open a pull request.

License
-------
[![License: MIT](https://img.shields.io/badge/License-mit-brightgreen.svg)](https://opensource.org/licenses/MIT)
