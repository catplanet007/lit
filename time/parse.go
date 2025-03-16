package time

import (
	"time"
)

func Parse(s string) (time.Time, error) {
	return ParseDateTime(s)
}

func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", s, DefaultLocation)
}

func ParseDateTime(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, DefaultLocation)
}

func ParseDateTimeMilli(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05.000", s, DefaultLocation)
}
