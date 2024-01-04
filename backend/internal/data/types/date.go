// Package types provides custom types for the application.
package types

import (
	"errors"
	"strings"
	"time"
)

// Date is a custom type that implements the MarshalJSON interface
// that applies date only formatting to the time.Time fields in order
// to avoid common time and timezone pitfalls when working with Times.
//
// Examples:
//
//	"2019-01-01" -> time.Time{2019-01-01 00:00:00 +0000 UTC}
//	"2019-01-01T21:10:30Z" -> time.Time{2019-01-01 00:00:00 +0000 UTC}
//	"2019-01-01T21:10:30+01:00" -> time.Time{2019-01-01 00:00:00 +0000 UTC}
type Date time.Time

func (d Date) Time() time.Time {
	return time.Time(d)
}

// DateFromTime returns a Date type from a time.Time type by stripping
// the time and timezone information.
func DateFromTime(t time.Time) Date {
	dateOnlyTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	return Date(dateOnlyTime)
}

// DateFromString returns a Date type from a string by parsing the
// string into a time.Time type and then stripping the time and
// timezone information.
//
// Errors are ignored and an empty Date is returned.
func DateFromString(s string) Date {
	if s == "" {
		return Date{}
	}

	try := [...]string{
		"2006-01-02",
		"01/02/2006",
		"2006/01/02",
		time.RFC3339,
	}

	for _, format := range try {
		t, err := time.Parse(format, s)
		if err == nil {
			return DateFromTime(t)
		}
	}

	return Date{}
}

func (d Date) String() string {
	if time.Time(d).IsZero() {
		return ""
	}

	return time.Time(d).Format("2006-01-02")
}

func (d Date) MarshalJSON() ([]byte, error) {
	if time.Time(d).IsZero() {
		return []byte(`""`), nil
	}

	return []byte(`"` + d.String() + `"`), nil
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	// unescape the string if necessary `\"` -> `"`
	str := strings.Trim(string(data), "\"")
	if str == "" || str == "null" || str == `""` {
		*d = Date{}
		return nil
	}

	try := [...]string{
		"2006-01-02",
		"01/02/2006",
		time.RFC3339,
	}

	set := false
	var t time.Time

	for _, format := range try {
		t, err = time.Parse(format, str)
		if err == nil {
			set = true
			break
		}
	}

	if !set {
		return errors.New("invalid date format")
	}

	// strip the time and timezone information
	*d = DateFromTime(t)

	return nil
}
