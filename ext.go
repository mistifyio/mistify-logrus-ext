package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

// DefaultSetup sets the logrus formatter and log level
func DefaultSetup(logLevel string) error {
	err := SetLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetFormatter(&MistifyFormatter{})
	return nil
}

// SetLevel parses and sets the log level
func SetLevel(logLevel string) error {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
