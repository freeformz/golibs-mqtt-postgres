package log

import (
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/mattn/go-colorable"
)

type logger struct {
	*logrus.Logger
}

var (
	rootLogger = newRootLogger()
)

func newRootLogger() *logger {
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
func IsDebugEnabled() bool {
	return rootLogger.Level >= logrus.DebugLevel
}

func IsInfoEnabled() bool {
	return rootLogger.Level >= logrus.InfoLevel
}

func IsWarnEnabled() bool {
	return rootLogger.Level >= logrus.WarnLevel
}

func IsErrorEnabled() bool {
	return rootLogger.Level >= logrus.ErrorLevel
}

func Debugf(format string, args ...interface{}) {
	if IsDebugEnabled() {
		rootLogger.Debugf(format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	rootLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	rootLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	rootLogger.Errorf(format, args...)
}

func Debug(args ...interface{}) {
	if IsDebugEnabled() {
		rootLogger.Debug(args...)
	}
}

func Info(args ...interface{}) {
	rootLogger.Info(args...)
}

func Warn(args ...interface{}) {
	rootLogger.Warn(args...)
}

func Error(args ...interface{}) {
	rootLogger.Error(args...)
}
