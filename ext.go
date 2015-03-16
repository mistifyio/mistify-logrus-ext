package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

func DefaultSetup(logLevel string) error {
	err := SetLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetFormatter(&MistifyFormatter{})
	return nil
}

func SetLevel(logLevel string) error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
