package types

import "time"

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

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		// TODO: Remove - used by legacy importer
		t, err = time.Parse("01/02/2006", s)

		if err != nil {
			return Date{}
		}
	}

	return DateFromTime(t)
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

func (d *Date) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == `""` {
		*d = Date{}
		return nil
	}

	// Try YYYY-MM-DD format
	var t time.Time
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		// Try default interface
		err = t.UnmarshalJSON(data)
		if err != nil {
			return err
		}
	}

	// strip the time and timezone information
	*d = DateFromTime(t)

	return nil
}
