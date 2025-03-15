package time

import (
	"log"
	"time"
)

var loc8 *time.Location

func init() {
	var err error
	loc8, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Printf("LoadLocation err: %v", err)
	}
}

func FormatDateTime(t time.Time) string {
	return t.In(loc8).Format("2006-01-02 15:04:05")
}

func FormatDate(t time.Time) string {
	return t.In(loc8).Format("2006-01-02")
}

func FormatTime(t time.Time) string {
	return t.In(loc8).Format("15:04:05")
}

func FormatDateTimeMilli(t time.Time) string {
	return t.In(loc8).Format("2006-01-02 15:04:05.000")
}

func FormatTimeMilli(t time.Time) string {
	return t.In(loc8).Format("15:04:05.000")
}
