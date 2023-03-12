package dlog

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

const correlationID = "1-581cf771-a006649127e371903a2de979"

func TestLogger(t *testing.T) {
	var buffer bytes.Buffer
	var fields logrus.Fields

	logger := NewLogger(nil)
	logger.Logger.Out = &buffer
	logger.Info("Hello World")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)

	assert.Equal(t, "Hello World", fields["message"])
	assert.Equal(t, "info", fields["level"])
	assert.NotNil(t, fields["timestamp"])
	assert.Nil(t, fields[CorrelationID])
}

func TestContextLogger(t *testing.T) {
	t.Run("should not include empty values", func(t *testing.T) {
		var buffer bytes.Buffer
		var fields logrus.Fields

		logger := NewLogger(&Config{
			AppName:       "",
			CorrelationID: "",
			Span:          "",
			Trace:         "",
			Parent:        "",
			Version:       "",
			Commit:        "",
			Build:         "",
			ReportCaller:  false,
		})
		logger.Logger.Out = &buffer
		logger.Info("Hello World")

		err := json.Unmarshal(buffer.Bytes(), &fields)
		assert.Nil(t, err)

		assert.Equal(t, "Hello World", fields["message"])
		assert.Equal(t, "info", fields["level"])
		assert.Nil(t, fields[CorrelationID])
		assert.Nil(t, fields[AppName])
		assert.Nil(t, fields[Span])
		assert.Nil(t, fields[Trace])
		assert.Nil(t, fields[Parent])
		assert.Nil(t, fields[Version])
		assert.Nil(t, fields[Commit])
		assert.Nil(t, fields[Build])
		assert.Nil(t, fields[Func])
		assert.Nil(t, fields[File])
		assert.NotNil(t, fields["timestamp"])
	})

	t.Run("Should include all", func(t *testing.T) {
		var buffer bytes.Buffer
		var fields logrus.Fields

		logger := NewLogger(&Config{
			AppName:       "MyService",
			CorrelationID: correlationID,
			Span:          "span-id",
			Trace:         "trace-id",
			Parent:        "parent-id",
			Version:       "version",
			Commit:        "commit",
			Build:         "build",
			ReportCaller:  true,
			Level:         "debug",
		})
		logger.Logger.Out = &buffer
		logger.Info("Hello World")

		err := json.Unmarshal(buffer.Bytes(), &fields)
		assert.Nil(t, err)

		assert.Equal(t, "Hello World", fields["message"])
		assert.Equal(t, "info", fields["level"])
		assert.Equal(t, correlationID, fields[CorrelationID])
		assert.Equal(t, "MyService", fields[AppName])
		assert.Equal(t, "span-id", fields[Span])
		assert.Equal(t, "trace-id", fields[Trace])
		assert.Equal(t, "parent-id", fields[Parent])
		assert.Equal(t, "version", fields[Version])
		assert.Equal(t, "commit", fields[Commit])
		assert.Equal(t, "build", fields[Build])
		assert.NotNil(t, fields["timestamp"])
		assert.Contains(t, fields[Func], "TestContextLogger")
		assert.Contains(t, fields[File], "main_test.go:84")
	})

	t.Run("should use text formatter", func(t *testing.T) {
		var buffer bytes.Buffer
		logger := NewLogger(&Config{
			AppName:       "myservice",
			CorrelationID: "",
			Span:          "",
			Trace:         "",
			Parent:        "",
			Version:       "",
			Commit:        "",
			Build:         "",
			ReportCaller:  false,
			Level:         "debug",
			Formatter:     "text", // default is json
		})
		logger.Logger.Out = &buffer
		logger.Info("Hello World")

		assert.Contains(t, buffer.String(), "INFO   ")
		assert.Contains(t, buffer.String(), " Hello World")
	})
}
