package logrusx_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"

	log "github.com/Sirupsen/logrus"
	h "github.com/bakins/test-helpers"
	logx "github.com/mistifyio/mistify-logrus-ext"
)

func TestErrorNotLost(t *testing.T) {
	var buffer bytes.Buffer
	logger := log.New()
	logger.Out = &buffer
	logger.Formatter = new(log.JSONFormatter)
	logger.Hooks.Add(new(logx.ErrorStructHook))

	logger.WithField("error", errors.New("test error message")).Info("test info message")

	entry := make(map[string]interface{})
	err := json.Unmarshal(buffer.Bytes(), &entry)
	h.Ok(t, err)

	errmap := entry["error"].(map[string]interface{})
	h.Equals(t, errmap["Message"], "test error message")
}

func TestErrorNotLostOnFieldNotNamedError(t *testing.T) {
	var buffer bytes.Buffer
	logger := log.New()
	logger.Out = &buffer
	logger.Formatter = new(log.JSONFormatter)
	logger.Hooks.Add(new(logx.ErrorStructHook))

	logger.WithField("errorish", errors.New("test error message")).Info("test info message")

	entry := make(map[string]interface{})
	err := json.Unmarshal(buffer.Bytes(), &entry)
	h.Ok(t, err)

	errmap := entry["errorish"].(map[string]interface{})
	h.Equals(t, errmap["Message"], "test error message")
}
