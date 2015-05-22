package sql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringType(t *testing.T) {
	var v StringType
	v = String(nil)
	assert.Equal(t, false, v.IsValid())
	assert.Equal(t, "", v.Get())

	v = String("")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "", v.Get())

	v = String("abcdefg")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "abcdefg", v.Get())

	v = String(int32(123))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "123", v.Get())

	v = String(int64(123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "123456789", v.Get())

	v = String(float64(132.123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "132.123456789", v.Get())
}

func TestNumericType(t *testing.T) {
	var v NumericType
	v = Numeric(nil)
	assert.Equal(t, false, v.IsValid())
	assert.Equal(t, float64(0), v.Get())

	v = Numeric("")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, float64(0), v.Get())

	v = Numeric("132.123456789")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, float64(132.123456789), v.Get())

	v = Numeric(int32(123))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, float64(123), v.Get())

	v = Numeric(int64(123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, float64(123456789), v.Get())

	v = Numeric(float64(132.123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, float64(132.123456789), v.Get())
}

func TestIntType(t *testing.T) {
	var v IntType
	v = Int(nil)
	assert.Equal(t, false, v.IsValid())
	assert.Equal(t, int32(0), v.Get())

	v = Int("")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int32(0), v.Get())

	v = Int("132.523456789")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int32(133), v.Get())

	v = Int(int32(123))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int32(123), v.Get())

	v = Int(int64(123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int32(123456789), v.Get())

	v = Int(float64(132.123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int32(132), v.Get())
}

func TestBigIntType(t *testing.T) {
	var v BigIntType
	v = BigInt(nil)
	assert.Equal(t, false, v.IsValid())
	assert.Equal(t, int64(0), v.Get())

	v = BigInt("")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int64(0), v.Get())

	v = BigInt("132.523456789")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int64(133), v.Get())

	v = BigInt(int32(123))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int64(123), v.Get())

	v = BigInt(int64(123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int64(123456789), v.Get())

	v = BigInt(float64(132.123456789))
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, int64(132), v.Get())
}

func TestTimestampType(t *testing.T) {
	var v TimestampType
	v = Timestamp(nil)
	assert.Equal(t, false, v.IsValid())
	assert.Equal(t, TimestampType{}.val, v.Get())

	v = Timestamp("2015-04-30T21:00:00")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "2015-05-01 06:00:00 +0900 JST", v.Get().String())
}

func TestSetStringTypeValue(t *testing.T) {
	var err error
	var v StringType
	var dVal interface{}
	v = StringType{}
	v.Set(nil)
	assert.Equal(t, false, v.IsValid())
	assert.Equal(t, "", v.Get())
	dVal, err = v.Value()
	assert.Nil(t, err)
	assert.Equal(t, nil, dVal)

	v = StringType{}
	v.Set("")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "", v.Get())
	dVal, err = v.Value()
	assert.Nil(t, err)
	assert.Equal(t, "", dVal)

	v = StringType{}
	v.Set("abcdefg")
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "abcdefg", v.Get())
	dVal, err = v.Value()
	assert.Nil(t, err)
	assert.Equal(t, "abcdefg", dVal)
}

func TestUnmarshalStringType(t *testing.T) {
	var err error
	var v StringType

	v = StringType{}
	err = v.UnmarshalJSON(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v.IsValid())

	v = StringType{}
	err = v.UnmarshalJSON([]byte(""))
	assert.Nil(t, err)
	assert.Equal(t, true, v.IsValid())

	v = StringType{}
	err = v.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err)
	assert.Equal(t, false, v.IsValid())

	v = StringType{}
	err = v.UnmarshalJSON([]byte("abcdefg"))
	assert.Nil(t, err)
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "abcdefg", v.Get())

	v = StringType{}
	err = v.UnmarshalJSON([]byte(`"abcdefg"`))
	assert.Nil(t, err)
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "abcdefg", v.Get())
}

func TestUnmarshalTimestampType(t *testing.T) {
	var err error
	var v TimestampType

	v = TimestampType{}
	err = v.UnmarshalJSON(nil)
	assert.Nil(t, err)
	assert.Equal(t, false, v.IsValid())

	v = TimestampType{}
	err = v.UnmarshalJSON([]byte("null"))
	assert.Nil(t, err)
	assert.Equal(t, false, v.IsValid())

	v = TimestampType{}
	err = v.UnmarshalJSON([]byte("2015-04-30T21:00:00"))
	assert.Nil(t, err)
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "2015-05-01 06:00:00 +0900 JST", v.Get().String())

	// golo-API dateformat
	v = TimestampType{}
	err = v.UnmarshalJSON([]byte("2015-04-30+21:00:00"))
	assert.Nil(t, err)
	assert.Equal(t, true, v.IsValid())
	assert.Equal(t, "2015-05-01 06:00:00 +0900 JST", v.Get().String())

	v = TimestampType{}
	err = v.UnmarshalJSON([]byte(`""`))
	assert.Nil(t, err)
	assert.Equal(t, false, v.IsValid())

	v = TimestampType{}
	err = v.UnmarshalJSON([]byte(`"0000-00-00+00:00:00"`))
	assert.Nil(t, err)
	assert.Equal(t, false, v.IsValid())
}
