package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

type ErrorMessageHook struct{}

func (hook *ErrorMessageHook) Fire(entry *log.Entry) error {
	for k, v := range entry.Data {
		if err, ok := v.(error); ok {
			entry.Data[k+".msg"] = err.Error()
		}
	}
	return nil
}

func (hook *ErrorMessageHook) Levels() []log.Level {
	return []log.Level{
		log.DebugLevel,
		log.InfoLevel,
		log.WarnLevel,
		log.ErrorLevel,
		log.FatalLevel,
		log.PanicLevel,
	}
}
