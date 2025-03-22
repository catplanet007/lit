package ltime

import (
	"time"
)

func Format(t time.Time) string {
	return FormatDateTime(t)
}

func FormatDateTime(t time.Time) string {
	return t.In(DefaultLocation).Format("2006-01-02 15:04:05")
}

func FormatDate(t time.Time) string {
	return t.In(DefaultLocation).Format("2006-01-02")
}

func FormatTime(t time.Time) string {
	return t.In(DefaultLocation).Format("15:04:05")
}

func FormatDateTimeMilli(t time.Time) string {
	return t.In(DefaultLocation).Format("2006-01-02 15:04:05.000")
}

func FormatTimeMilli(t time.Time) string {
	return t.In(DefaultLocation).Format("15:04:05.000")
}
