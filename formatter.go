package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

type (
	MistifyFormatter struct {
		log.JSONFormatter
	}
	FieldError struct {
		Error   error
		Message string
	}
)

func (f *MistifyFormatter) Format(entry *log.Entry) ([]byte, error) {
	for k, v := range entry.Data {
		if err, ok := v.(error); ok {
			entry.Data[k] = FieldError{err, err.Error()}
		}
	}
	return f.JSONFormatter.Format(entry)
}
