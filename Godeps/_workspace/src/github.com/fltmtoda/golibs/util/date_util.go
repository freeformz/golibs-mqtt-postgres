package util

import (
	"fmt"
	"strings"
	"time"
)

const (
	LOCATION_JST = "JST"
	LOCATION_UTC = "UTC"
	LOCATION_MST = "MST"

	FORMAT_YYYYMMDDHHMMSS_WITH_HYPHEN = "2006-01-02 15:04:05"
)

func CurrentJSTTime() time.Time {
	return time.Now()
}

func CurrentUTCTime() time.Time {
	t := time.Now()
	if isJST(&t) {
		return t.UTC()
	}
	return t
}

func ToUTC(t time.Time) time.Time {
	if isJST(&t) {
		return t.UTC()
	}
	return t
}

func ToJST(t time.Time) time.Time {
	if isJST(&t) {
		return t
	}
	return t.Local()
}

func ParseJstTime(yyyyMMddHHmmss string) (time.Time, error) {
	sTime := yyyyMMddHHmmss
	if len(sTime) > 19 {
		sTime = sTime[0:19]
	}
	t, err := time.Parse(
		FORMAT_YYYYMMDDHHMMSS_WITH_HYPHEN+HALF_SPACE+LOCATION_MST,
		sTime+".000000"+HALF_SPACE+LOCATION_JST,
	)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ParseUtcTime(yyyyMMddHHmmss string) (time.Time, error) {
	sTime := yyyyMMddHHmmss
	if len(sTime) > 19 {
		sTime = sTime[0:19]
	}
	t, err := time.Parse(
		FORMAT_YYYYMMDDHHMMSS_WITH_HYPHEN,
		sTime+".000000",
	)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func CheckSameMonth(from, to time.Time) bool {
	d1 := ToJST(from)
	d2 := ToJST(to)
	return d1.Year() == d2.Year() && d1.Month().String() == d2.Month().String()
}
func CheckSameDate(from, to time.Time) bool {
	d1 := ToJST(from)
	d2 := ToJST(to)
	return d1.Year() == d2.Year() && d1.Month().String() == d2.Month().String() && d1.Day() == d2.Day()
}
func ToFirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

/**************************
 internal function
**************************/
func isJST(j *time.Time) bool {
	return strings.LastIndex(fmt.Sprintf("%v", j), LOCATION_JST) != -1
}
func isUTC(j *time.Time) bool {
	return strings.LastIndex(fmt.Sprintf("%v", j), LOCATION_UTC) != -1
}
