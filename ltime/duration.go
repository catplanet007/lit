package ltime

import (
	"fmt"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	if len(b) < 2 {
		return fmt.Errorf("parse TimeDuration [%s] invalid", b)
	}
	t, err := time.ParseDuration(string(b[1 : len(b)-1]))
	*d = Duration(t)
	return err
}

func (d Duration) D() time.Duration {
	return time.Duration(d)
}
