package env

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetenvStringValue(t *testing.T) {
	var env *Env
	key, value := "STRING_KEY", "ABCDEFG"

	env = GetEnv(key, "")
	assert.Equal(t,
		"",
		env.String(),
	)

	os.Setenv(key, value)
	env = GetEnv(key, nil)
	assert.Equal(t,
		value,
		env.String(),
	)

	os.Remove(key)
}

func TestGetenvFloat64Value(t *testing.T) {
	var env *Env
	key, value := "FLAOT64_KEY", float64(123456.123)

	env = GetEnv(key, nil)
	assert.Equal(t,
		float64(0),
		env.Float64(),
	)

	os.Setenv(key, fmt.Sprintf("%v", value))
	env = GetEnv(key, nil)
	assert.Equal(t,
		value,
		env.Float64(),
	)

	os.Remove(key)
}

func TestGetenvUint64Value(t *testing.T) {
	var env *Env
	key, value := "Uint64_KEY", uint64(123456)

	env = GetEnv(key, nil)
	assert.Equal(t,
		uint64(0),
		env.Uint64(),
	)

	os.Setenv(key, fmt.Sprintf("%v", value))
	env = GetEnv(key, nil)
	assert.Equal(t,
		value,
		env.Uint64(),
	)

	os.Remove(key)
}

func TestGetenvInt64Value(t *testing.T) {
	var env *Env
	key, value := "INT64_KEY", int64(-123456)

	env = GetEnv(key, nil)
	assert.Equal(t,
		int64(0),
		env.Int64(),
	)

	os.Setenv(key, fmt.Sprintf("%v", value))
	env = GetEnv(key, nil)
	assert.Equal(t,
		value,
		env.Int64(),
	)

	os.Remove(key)
}
