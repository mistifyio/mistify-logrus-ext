package logrusx

import (
	log "github.com/Sirupsen/logrus"
)

func DefaultSetup(logLevel string) error {
	log.SetFormatter(&ExtJSONFormatter{})
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(level)
	return nil
}
