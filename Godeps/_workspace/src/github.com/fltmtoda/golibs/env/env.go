package env

import (
	"os"

	"github.com/fltmtoda/golibs/util"
)

type Env struct {
	val    string
	defVal interface{}
}

func GetEnv(key string, defaultValue interface{}) *Env {
	return &Env{
		val:    os.Getenv(key),
		defVal: defaultValue,
	}
}

func (e *Env) String() string {
	if e.hasValue() {
		return e.val
	}
	return util.ToString(e.defVal)
}

func (e *Env) Float64() float64 {
	if e.hasValue() {
		return util.ToFloat64(e.val)
	}
	return util.ToFloat64(e.defVal)
}

func (e *Env) Uint64() uint64 {
	if e.hasValue() {
		return util.ToUint64(e.val)
	}
	return util.ToUint64(e.defVal)
}

func (e *Env) Int64() int64 {
	if e.hasValue() {
		return util.ToInt64(e.val)
	}
	return util.ToInt64(e.defVal)
}

func (e *Env) hasValue() bool {
	return e.val != util.BLANK
}
