package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	logx "github.com/mistifyio/mistify-logrus-ext"
)

type TestError struct {
	Val1 string
	Val2 string
}

func (e *TestError) Error() string {
	return "This is a complex error"
}

func main() {
	fmt.Println("This is a test of logrusx")
	log.Info("Logging before setup")
	log.WithFields(log.Fields{
		"foo":    "bar",
		"error1": fmt.Errorf("This is a basic error"),
		"error2": &TestError{"foobar", "baz"},
	}).Info("Logging before setup with fields")

	err1 := logx.DefaultSetup("bad log level")
	if err1 != nil {
		log.WithFields(log.Fields{
			"error": err1,
			"func":  "logx.DefaultSetup",
		}).Error("Could not do logrusx default setup with bad log level")
	}

	err2 := logx.DefaultSetup("debug")
	if err2 != nil {
		log.WithFields(log.Fields{
			"error": err2,
			"func":  "logx.DefaultSetup",
		}).Error("Could not do logrusx default setup with good log level")
	}

	log.Info("Logging after setup")
	log.WithFields(log.Fields{
		"foo":    "bar",
		"error1": fmt.Errorf("This is a basic error"),
		"error2": &TestError{"foobar", "baz"},
	}).Info("Logging after setup with fields")
}
