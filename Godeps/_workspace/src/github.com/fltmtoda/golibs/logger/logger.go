package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
	"os"
	"strings"
)

type logger struct {
	internal *logrus.Logger
}

var instance *logger

func GetLogger() *logger {
	if instance == nil {
		instance = newLogger()
	}
	return instance
}

func newLogger() *logger {
	switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn", "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(newFormatter(true))
	return &logger{logrus.StandardLogger()}
}

/**************************
 logger function
**************************/
func (l *logger) IsDebugEnabled() bool {
	return l.internal.Level >= logrus.DebugLevel
}

func (l *logger) IsInfoEnabled() bool {
	return l.internal.Level >= logrus.InfoLevel
}

func (l *logger) IsWarnEnabled() bool {
	return l.internal.Level >= logrus.WarnLevel
}

func (l *logger) IsErrorEnabled() bool {
	return l.internal.Level >= logrus.ErrorLevel
}

func (l *logger) Debugf(format string, args ...interface{}) {
	if l.IsDebugEnabled() {
		l.internal.Debugf(format, args...)
	}
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.internal.Infof(format, args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.internal.Warnf(format, args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.internal.Errorf(format, args...)
}

func (l *logger) Debug(args ...interface{}) {
	if l.IsDebugEnabled() {
		l.internal.Debug(args...)
	}
}

func (l *logger) Info(args ...interface{}) {
	l.internal.Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.internal.Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.internal.Error(args...)
}
