package logrusx_test

import (
	"encoding/json"
	"errors"
	"testing"

	log "github.com/Sirupsen/logrus"
	logx "github.com/mistifyio/mistify-logrus-ext"
)

func TestErrorNotLost(t *testing.T) {
	formatter := &logx.ExtJSONFormatter{}

	msg := "test error message"
	b, err := formatter.Format(log.WithField("error", errors.New(msg)))
	if err != nil {
		t.Fatal("Could not format log entry: ", err)
	}

	entry := make(map[string]interface{})
	err = json.Unmarshal(b, &entry)
	if err != nil {
		t.Fatal("Could not unmarshal formatted log entry: ", err)
	}

	errmap := entry["error"].(map[string]interface{})
	if errmap["Message"] != "test error message" {
		t.Fatal("Error message field not added")
	}
}

func TestErrorNotLostOnFieldNotNamedError(t *testing.T) {
	formatter := &logx.ExtJSONFormatter{}

	msg := "test errorish message"
	b, err := formatter.Format(log.WithField("errorish", errors.New(msg)))
	if err != nil {
		t.Fatal("Could not format log entry: ", err)
	}

	entry := make(map[string]interface{})
	err = json.Unmarshal(b, &entry)
	if err != nil {
		t.Fatal("Could not unmarshal formatted log entry: ", err)
	}

	errmap := entry["errorish"].(map[string]interface{})
	if errmap["Message"] != msg {
		t.Fatal("Error message field not set")
	}
}
