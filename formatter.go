// Package logrusx is a logrus formatter that adds better error value handling
// to the logrus.JSONFormatter
package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

type (
	// MistifyFormatter is a custom logrus formatter extending JSONFormatter
	MistifyFormatter struct {
		log.JSONFormatter
	}

	// FieldError contains both the error struct and error message as explicit
	// properties, including both when JSON marshaling.
	FieldError struct {
		Error   error
		Message string
	}
)

// Format replaces any error field values with a FieldError and produces a JSON
// formatted log entry
func (f *MistifyFormatter) Format(entry *log.Entry) ([]byte, error) {
	for k, v := range entry.Data {
		if err, ok := v.(error); ok {
			entry.Data[k] = FieldError{err, err.Error()}
		}
	}
	return f.JSONFormatter.Format(entry)
}
