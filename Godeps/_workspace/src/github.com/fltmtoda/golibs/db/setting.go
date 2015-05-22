package db

import (
	"strings"
)

type Setting struct {
	Url          string
	MaxIdleConns int
	MaxOpenConns int
}

/**************************
 Setting function
**************************/
func (s *Setting) IsDefaultSetting() bool {
	return s.MaxIdleConns == 0 && s.MaxOpenConns == 0
}

func (s *Setting) Dialect() string {
	if strings.HasPrefix(s.Url, "postgres") {
		return "postgres"
	}
	return "none dialect"
}
