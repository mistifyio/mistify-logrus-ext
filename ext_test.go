package logrusx_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	log "github.com/Sirupsen/logrus"
	h "github.com/bakins/test-helpers"
	logx "github.com/mistifyio/mistify-logrus-ext"
)

func TestSetLevel(t *testing.T) {
	var err error
	defer log.SetLevel(log.InfoLevel)
	err = logx.SetLevel("foobar")
	h.Assert(t, err != nil, "expected log level error")

	err = logx.SetLevel("debug")
	h.Ok(t, err)

	std := log.StandardLogger()
	h.Equals(t, log.DebugLevel, std.Level)

}

func TestDefaultSetup(t *testing.T) {
	var err error
	defer log.SetLevel(log.InfoLevel)
	defer log.SetFormatter(&log.TextFormatter{})
	err = logx.DefaultSetup("foobar")
	h.Assert(t, err != nil, "expected log level error")

	err = logx.DefaultSetup("debug")
	h.Ok(t, err)

	std := log.StandardLogger()
	_, ok := std.Formatter.(*logx.MistifyFormatter)
	h.Assert(t, ok, "expected MistifyFormatter to be set")
}

func TestLogReturnedErr(t *testing.T) {
	var buffer bytes.Buffer
	var out string

	log.SetOutput(&buffer)

	fn := func() error {
		return nil
	}

	logx.LogReturnedErr(fn, nil, "qwerty")
	out = buffer.String()
	h.Assert(t, !strings.Contains(out, "foobar"), "did not expect error logged")
	h.Assert(t, !strings.Contains(out, "qwerty"), "did not expect message logged")
	buffer.Truncate(0)

	fnE := func() error {
		return errors.New("foobar")
	}

	logx.LogReturnedErr(fnE, nil, "qwerty")
	out = buffer.String()
	h.Assert(t, strings.Contains(out, "foobar"), "expected error logged")
	h.Assert(t, strings.Contains(out, "qwerty"), "expected message logged")
}
