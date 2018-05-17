package dlog

import (
	"testing"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/sirupsen/logrus"
	"encoding/json"
)

const requestId = "1-581cf771-a006649127e371903a2de979"

func TestLogger(t *testing.T) {
	var buffer bytes.Buffer
	var fields logrus.Fields

	logger := NewLogger("MyService")
	logger.Logger.Out = &buffer
	logger.Info("Hello World")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)

	assert.Equal(t, fields["message"], "Hello World")
	assert.Equal(t, fields["level"], "info")
	assert.Equal(t, fields[Service], "MyService")
	assert.NotNil(t, fields["timestamp"])
	assert.Nil(t, fields[RequestIO])
}

func TestContextLogger(t *testing.T) {
	var buffer bytes.Buffer
	var fields logrus.Fields

	logger := NewRequestLogger(requestId, "MyService")
	logger.Logger.Out = &buffer
	logger.Info("Hello World")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	assert.Nil(t, err)

	assert.Equal(t, fields["message"], "Hello World")
	assert.Equal(t, fields["level"], "info")
	assert.Equal(t, fields[RequestIO], requestId)
	assert.Equal(t, fields[Service], "MyService")
	assert.NotNil(t, fields["timestamp"])
}
