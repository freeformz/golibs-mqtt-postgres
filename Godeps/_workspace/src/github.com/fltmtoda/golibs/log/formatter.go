package log

import (
	"bytes"
	"fmt"

	"github.com/Sirupsen/logrus"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
	gray    = 39
)

var (
	formatParams = map[logrus.Level]formatParam{
		logrus.DebugLevel: formatParam{
			color:     blue,
			levelText: "DEBUG",
		},
		logrus.InfoLevel: formatParam{
			color:     green,
			levelText: "INFO",
		},
		logrus.WarnLevel: formatParam{
			color:     yellow,
			levelText: "WARN",
		},
		logrus.ErrorLevel: formatParam{
			color:     red,
			levelText: "ERROR",
		},
	}
)

type (
	formatter struct {
		logrus.TextFormatter
	}

	formatParam struct {
		color     int
		levelText string
	}
)

func newFormatter(color bool) *formatter {
	return &formatter{
		logrus.TextFormatter{
			ForceColors:   color,
			DisableColors: !color,
		},
	}
}

/**************************
 formatter function
**************************/
func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	formatParams := formatParams[entry.Level]
	b := &bytes.Buffer{}
	if f.ForceColors && !f.DisableColors {
		fmt.Fprintf(b,
			"\x1b[%dm[%-5s]\x1b[0m %s",
			formatParams.color,
			formatParams.levelText,
			entry.Message,
		)
	} else {
		fmt.Fprintf(b,
			"[%-5s] %s",
			formatParams.levelText,
			entry.Message,
		)
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
