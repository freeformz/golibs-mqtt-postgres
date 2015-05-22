package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurrentTime(t *testing.T) {
	actualJst := CurrentJSTTime()
	actualUtc := CurrentUTCTime()
	assert.NotEqual(t,
		actualJst,
		actualUtc,
	)
	assert.Equal(t,
		actualJst.Format("2006-01-02 15:04"),
		ToJST(actualUtc).Format("2006-01-02 15:04"),
	)
	assert.Equal(t,
		actualUtc.Format("2006-01-02 15:04"),
		ToUTC(actualJst).Format("2006-01-02 15:04"),
	)
}

func TestJstTime(t *testing.T) {
	sDate := "2015-01-02 12:34:56"
	actualJst, err := ParseJstTime(sDate)
	assert.Nil(t, err)
	assert.Equal(t,
		sDate+" +0900 JST",
		fmt.Sprintf("%v", actualJst),
	)
	assert.True(t, isJST(&actualJst))
	assert.False(t, isUTC(&actualJst))

	actualUtc := ToUTC(actualJst)
	assert.False(t, isJST(&actualUtc))
	assert.True(t, isUTC(&actualUtc))
	assert.Equal(t,
		"2015-01-02 03:34:56 +0000 UTC",
		fmt.Sprintf("%v", actualUtc),
	)
}

func TestParseUtcTime(t *testing.T) {
	sDate := "2015-01-02 12:34:56"
	actualUtc, err := ParseUtcTime(sDate)
	assert.Nil(t, err)
	assert.Equal(t,
		sDate+" +0000 UTC",
		fmt.Sprintf("%v", actualUtc),
	)
	assert.False(t, isJST(&actualUtc))
	assert.True(t, isUTC(&actualUtc))

	actualJst := ToJST(actualUtc)
	assert.True(t, isJST(&actualJst))
	assert.False(t, isUTC(&actualJst))
	assert.Equal(t,
		"2015-01-02 21:34:56 +0900 JST",
		fmt.Sprintf("%v", actualJst),
	)
}

func TestCheckSameMonth(t *testing.T) {
	d1, err := ParseJstTime("2015-01-02 00:00:00")
	assert.Nil(t, err)
	d2, err := ParseJstTime("2015-01-02 00:00:00")
	assert.Nil(t, err)
	d3, err := ParseJstTime("2015-02-01 00:00:00")
	assert.Nil(t, err)

	assert.True(t, CheckSameMonth(d1, d2))
	assert.False(t, CheckSameMonth(d1, d3))

	assert.True(t, CheckSameMonth(d1, ToUTC(d2)))
	assert.True(t, CheckSameMonth(ToUTC(d1), d2))
	assert.False(t, CheckSameDate(d1, ToUTC(d3)))
}

func TestCheckSameDate(t *testing.T) {
	d1, err := ParseJstTime("2015-01-02 12:34:56")
	assert.Nil(t, err)
	d2, err := ParseJstTime("2015-01-02 21:34:56")
	assert.Nil(t, err)
	d3, err := ParseJstTime("2015-01-03 00:00:00")
	assert.Nil(t, err)

	assert.True(t, CheckSameDate(d1, d2))
	assert.False(t, CheckSameDate(d1, d3))

	assert.True(t, CheckSameDate(d1, ToUTC(d2)))
	assert.True(t, CheckSameDate(ToUTC(d1), d2))
	assert.False(t, CheckSameDate(d1, ToUTC(d3)))
}
