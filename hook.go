package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

type (
	ErrorStructHook struct{}
	FieldError      struct {
		Error   error
		Message string
	}
)

func (hook *ErrorStructHook) Fire(entry *log.Entry) error {
	for k, v := range entry.Data {
		if err, ok := v.(error); ok {
			entry.Data[k] = FieldError{err, err.Error()}
		}
	}
	return nil
}

func (hook *ErrorStructHook) Levels() []log.Level {
	return []log.Level{
		log.DebugLevel,
		log.InfoLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	}
}
