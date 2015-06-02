package db

import (
	"strings"
)

type Setting struct {
	URL          string
	MaxIdleConns int
	MaxOpenConns int
}

/**************************
 Setting function
**************************/
func (s *Setting) Dialect() string {
	if strings.HasPrefix(s.URL, "postgres") {
		return "postgres"
	}
	return ""
}
