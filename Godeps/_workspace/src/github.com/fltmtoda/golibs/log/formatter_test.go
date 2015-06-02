package log

import (
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	message = "logger test."
)

func TestDebugLogFormat(t *testing.T) {
	assertFormatter(t,
		"[DEBUG] "+message+"\n",
		logrus.DebugLevel,
		message,
	)
}

func TestInfoLogFormat(t *testing.T) {
	assertFormatter(t,
		"[INFO ] "+message+"\n",
		logrus.InfoLevel,
		message,
	)
}

func TestWarnLogFormat(t *testing.T) {
	assertFormatter(t,
		"[WARN ] "+message+"\n",
		logrus.WarnLevel,
		message,
	)
}

func TestErrorLogFormat(t *testing.T) {
	assertFormatter(t,
		"[ERROR] "+message+"\n",
		logrus.ErrorLevel,
		message,
	)
}

func assertFormatter(
	t *testing.T,
	expected string,
	logLevel logrus.Level,
	msg string,
) {
	entry := logrus.NewEntry(logrus.StandardLogger())
	entry.Level = logLevel
	entry.Message = msg
	msgBytes, _ := newFormatter(false).Format(entry)
	assert.Equal(t,
		expected,
		string(msgBytes),
	)
}
