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
