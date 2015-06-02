package util

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/fltmtoda/golibs/log"
)

func ToString(val interface{}) string {
	if val == nil {
		return BLANK
	}
	switch v := val.(type) {
	case string:
		return v
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []byte:
		return ToString(string(v))
	default:
		return fmt.Sprintf("%+v", v)
	}
}

func ToInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case string:
		if v == BLANK {
			return 0
		}
		ret, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fVal, err := strconv.ParseFloat(v, 64)
			if err == nil {
				return ToInt64(fVal)
			}
			log.Warnf("Unable to convert string to int64 %v", v)
			return 0
		}
		return ret
	case float32, float64:
		return ToInt64(fmt.Sprintf("%.0f", v))
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case uintptr:
		return int64(v)
	case []byte:
		return ToInt64(string(v))
	default:
		log.Warnf("Unknown format? %v %s", val, reflect.ValueOf(v))
		return 0
	}
}

func ToUint64(val interface{}) uint64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case string:
		if v == BLANK {
			return 0
		}
		ret, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			fVal, err := strconv.ParseFloat(v, 64)
			if err == nil {
				return ToUint64(fVal)
			}
			log.Warnf("Unable to convert string to uint64 %v", v)
			return 0
		}
		return ret
	case float32, float64:
		return ToUint64(fmt.Sprintf("%.0f", v))
	case int:
		return uint64(v)
	case int8:
		return uint64(v)
	case int16:
		return uint64(v)
	case int32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case uintptr:
		return uint64(v)
	case []byte:
		return ToUint64(string(v))
	default:
		log.Warnf("Unknown format? %v %s", val, reflect.ValueOf(v))
		return 0
	}
}

func ToFloat64(val interface{}) float64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case string:
		if v == BLANK {
			return 0
		}
		ret, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Warnf("Unable to convert string to float64 %v", v)
			return 0
		}
		return ret
	case float32:
		return float64(v)
	case float64:
		return v
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case uintptr:
		return float64(v)
	case []byte:
		return ToFloat64(string(v))
	default:
		log.Warnf("Unknown format? %v %s", val, reflect.ValueOf(v))
		return 0
	}
}
