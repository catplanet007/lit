package time

import (
	"time"
)

var DefaultLocation *time.Location

func init() {
	var err error
	DefaultLocation, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
}
