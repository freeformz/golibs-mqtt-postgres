package log

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	defaultLogLevel := os.Getenv("LOG_LEVEL")
	v := m.Run()
	os.Setenv("LOG_LEVEL", defaultLogLevel)
	os.Exit(v)
}

func TestDebugLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	rootLogger = newRootLogger()

	assert.True(t, IsDebugEnabled())
	assert.True(t, IsInfoEnabled())
	assert.True(t, IsWarnEnabled())
	assert.True(t, IsErrorEnabled())
}

func TestInfoLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "info")
	rootLogger = newRootLogger()

	assert.False(t, IsDebugEnabled())
	assert.True(t, IsInfoEnabled())
	assert.True(t, IsWarnEnabled())
	assert.True(t, IsErrorEnabled())
}

func TestWarnLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "warn")
	rootLogger = newRootLogger()

	assert.False(t, IsDebugEnabled())
	assert.False(t, IsInfoEnabled())
	assert.True(t, IsWarnEnabled())
	assert.True(t, IsErrorEnabled())
}

func TestErrorLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "error")
	rootLogger = newRootLogger()

	assert.False(t, IsDebugEnabled())
	assert.False(t, IsInfoEnabled())
	assert.False(t, IsWarnEnabled())
	assert.True(t, IsErrorEnabled())
}
