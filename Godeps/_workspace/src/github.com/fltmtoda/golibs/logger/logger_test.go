package logger

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	v := m.Run()
	os.Exit(v)
}

func TestDebugLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	logger := newLogger()
	assert.True(t, logger.IsDebugEnabled())
	assert.True(t, logger.IsInfoEnabled())
	assert.True(t, logger.IsWarnEnabled())
	assert.True(t, logger.IsErrorEnabled())
}

func TestInfoLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "info")
	logger := newLogger()
	assert.False(t, logger.IsDebugEnabled())
	assert.True(t, logger.IsInfoEnabled())
	assert.True(t, logger.IsWarnEnabled())
	assert.True(t, logger.IsErrorEnabled())
}

func TestWarnLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "warn")
	logger := newLogger()
	assert.False(t, logger.IsDebugEnabled())
	assert.False(t, logger.IsInfoEnabled())
	assert.True(t, logger.IsWarnEnabled())
	assert.True(t, logger.IsErrorEnabled())
}

func TestErrorLog(t *testing.T) {
	os.Setenv("LOG_LEVEL", "error")
	logger := newLogger()
	logger.Debug("%s", "abc")
	assert.False(t, logger.IsDebugEnabled())
	assert.False(t, logger.IsInfoEnabled())
	assert.False(t, logger.IsWarnEnabled())
	assert.True(t, logger.IsErrorEnabled())
}
