package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

func DefaultSetup(logLevel string) error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	log.SetFormatter(&log.JSONFormatter{})
	log.AddHook(&ErrorStructHook{})
	return nil
}
