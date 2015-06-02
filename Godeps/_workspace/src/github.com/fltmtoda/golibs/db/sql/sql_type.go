package sql

import (
	"database/sql/driver"
	"strings"
	"time"

	"github.com/fltmtoda/golibs/log"
	"github.com/fltmtoda/golibs/util"
)

var (
	emptyTime = TimestampType{}.val
)

type (
	baseType struct {
		valid bool
		val   interface{}
	}

	StringType struct {
		val string
		baseType
	}

	NumericType struct {
		val float64
		baseType
	}

	IntType struct {
		val int32
		baseType
	}

	BigIntType struct {
		val int64
		baseType
	}

	TimestampType struct {
		val time.Time
		baseType
	}
)

func String(value interface{}) StringType {
	ret := StringType{}
	ret.Set(value)
	return ret
}

func Numeric(value interface{}) NumericType {
	ret := NumericType{}
	ret.Set(value)
	return ret
}

func Int(value interface{}) IntType {
	ret := IntType{}
	ret.Set(value)
	return ret
}

func BigInt(value interface{}) BigIntType {
	ret := BigIntType{}
	ret.Set(value)
	return ret
}

func Timestamp(value interface{}) TimestampType {
	ret := TimestampType{}
	ret.Set(value)
	return ret
}

func (t *baseType) Scan(value interface{}) error {
	t.Set(value)
	return nil
}
func (t *baseType) Set(value interface{}) {
	// override function
}
func (t *baseType) IsValid() bool {
	return t.valid
}

/**************************
 StringType function
**************************/
func (t *StringType) Scan(value interface{}) error {
	t.Set(value)
	return nil
}
func (t StringType) Value() (driver.Value, error) {
	if !t.valid {
		return nil, nil
	}
	return t.val, nil
}
func (t *StringType) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	sVal := string(data)
	if sVal == `""` || sVal == "null" {
		return nil
	}
	if strings.HasPrefix(sVal, "\"") && strings.HasPrefix(sVal, "\"") {
		sVal = strings.Replace(sVal, `\"`, `"`, -1)
		t.Set(sVal[1 : len(sVal)-1])
	} else {
		t.Set(sVal)
	}
	if log.IsDebugEnabled() {
		log.Debugf("UnmarshalJSON StringType: %v => %v", sVal, t)
	}
	return nil
}
func (t *StringType) Set(value interface{}) {
	if value != nil {
		t.valid, t.val = true, util.ToString(value)
	} else {
		t.valid, t.val = false, util.BLANK
	}
}
func (t *StringType) Get() string {
	if t.valid {
		return t.val
	}
	return util.BLANK
}

/**************************
 NumericType function
**************************/
func (t *NumericType) Scan(value interface{}) error {
	t.Set(value)
	return nil
}
func (t NumericType) Value() (driver.Value, error) {
	if !t.valid {
		return nil, nil
	}
	return util.ToString(t.val), nil
}
func (t *NumericType) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	sVal := string(data)
	if sVal == `""` || sVal == "null" {
		return nil
	}
	if strings.HasPrefix(sVal, "\"") && strings.HasPrefix(sVal, "\"") {
		t.Set(sVal[1 : len(sVal)-1])
	} else {
		t.Set(sVal)
	}
	if log.IsDebugEnabled() {
		log.Debugf("UnmarshalJSON NumericType: %v => %v", sVal, t)
	}
	return nil
}
func (t *NumericType) Set(value interface{}) {
	if value != nil {
		t.valid, t.val = true, util.ToFloat64(value)
	} else {
		t.valid, t.val = false, 0
	}
}
func (t *NumericType) Get() float64 {
	if t.valid {
		return t.val
	}
	return float64(0)
}

/**************************
 IntType function
**************************/
func (t *IntType) Scan(value interface{}) error {
	t.Set(value)
	return nil
}
func (t IntType) Value() (driver.Value, error) {
	if !t.valid {
		return nil, nil
	}
	return util.ToString(t.val), nil
}
func (t *IntType) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	sVal := string(data)
	if sVal == `""` || sVal == "null" {
		return nil
	}
	if strings.HasPrefix(sVal, "\"") && strings.HasPrefix(sVal, "\"") {
		t.Set(sVal[1 : len(sVal)-1])
	} else {
		t.Set(sVal)
	}

	if log.IsDebugEnabled() {
		log.Debugf("UnmarshalJSON IntType: %v => %v", sVal, t)
	}
	return nil
}
func (t *IntType) Set(value interface{}) {
	if value != nil {
		t.valid, t.val = true, int32(util.ToInt64(value))
	} else {
		t.valid, t.val = false, int32(0)
	}
}
func (t *IntType) Get() int32 {
	if t.valid {
		return t.val
	}
	return int32(0)
}

/**************************
 BigIntType function
**************************/
func (t *BigIntType) Scan(value interface{}) error {
	t.Set(value)
	return nil
}
func (t BigIntType) Value() (driver.Value, error) {
	if !t.valid {
		return nil, nil
	}
	return t.val, nil
}
func (t *BigIntType) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	sVal := string(data)
	if sVal == `""` || sVal == "null" {
		return nil
	}
	if strings.HasPrefix(sVal, "\"") && strings.HasPrefix(sVal, "\"") {
		t.Set(sVal[1 : len(sVal)-1])
	} else {
		t.Set(sVal)
	}
	if log.IsDebugEnabled() {
		log.Debugf("UnmarshalJSON BigIntType: %v => %v", sVal, t)
	}
	return nil
}
func (t *BigIntType) Set(value interface{}) {
	if value != nil {
		t.valid, t.val = true, util.ToInt64(value)
	} else {
		t.valid, t.val = false, 0
	}
}
func (t *BigIntType) Get() int64 {
	if t.valid {
		return t.val
	}
	return int64(0)
}

/**************************
 TimestampType function
**************************/
func (t *TimestampType) Scan(value interface{}) error {
	t.Set(value)
	return nil
}
func (t TimestampType) Value() (driver.Value, error) {
	if !t.valid {
		return nil, nil
	}
	return t.val, nil
}
func (t *TimestampType) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}
	sVal := string(data)
	if sVal == `""` || sVal == "null" || sVal == `"0000-00-00+00:00:00"` {
		return nil
	}
	sVal = strings.Replace(sVal, "\"", "", -1)
	sVal = strings.Replace(sVal, "+", " ", -1)
	if sVal != "" {
		if strings.Contains(sVal, "-") {
			t.Set(sVal)
		} else {
			t.Set(time.Unix(util.ToInt64(sVal), 0).Add(9 * time.Hour)) //TODO Long値で取得した場合,UTC=>JSTへの変換が必要
		}
	}
	if log.IsDebugEnabled() {
		log.Debugf("UnmarshalJSON TimestampType: %v => %v", sVal, t)
	}
	return nil
}
func (t *TimestampType) Set(value interface{}) {
	if value != nil {
		switch v := value.(type) {
		case time.Time:
			t.valid, t.val = true, v
		default:
			sTime := util.ToString(v)
			if sTime != util.BLANK && sTime != "null" {
				sTime = strings.Replace(sTime, "T", " ", -1)
				if len(sTime) > 26 {
					sTime = sTime[0:26]
				}
				sTime = strings.TrimSpace(sTime)
				parseTime, err := time.Parse("2006-01-02 15:04:05.999999", sTime)
				if err != nil {
					log.Errorf("Time parse error: %v", err)
					break
				}
				t.valid, t.val = true, parseTime.Local()
			}
		}
	}
	if !t.valid {
		t.val = emptyTime
	}
}
func (t *TimestampType) Get() time.Time {
	if t.valid {
		return t.val
	}
	return emptyTime
}
