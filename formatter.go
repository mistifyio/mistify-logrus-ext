package logrusx

import (
	"encoding/json"
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
)

type ExtJSONFormatter struct{}

type FieldError struct {
	Error   error
	Message string
}

func (f *ExtJSONFormatter) prefixFieldClashes(data log.Fields) {
	_, ok := data["time"]
	if ok {
		data["fields.time"] = data["time"]
	}

	_, ok = data["msg"]
	if ok {
		data["fields.msg"] = data["msg"]
	}

	_, ok = data["level"]
	if ok {
		data["fields.level"] = data["level"]
	}
}

func (f *ExtJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	data := make(log.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		if err, ok := v.(error); ok {
			data[k] = FieldError{err, err.Error()}
		} else {
			data[k] = v
		}
	}
	f.prefixFieldClashes(data)
	data["time"] = entry.Time.Format(time.RFC3339)
	data["msg"] = entry.Message
	data["level"] = entry.Level.String()

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
