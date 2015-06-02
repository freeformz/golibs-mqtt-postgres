package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettingDialect(t *testing.T) {
	s := Setting{
		URL: "postgres://user:pass@localhost:5432/postgres",
	}
	assert.Equal(t,
		"postgres",
		s.Dialect(),
	)
}
