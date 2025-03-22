package lconcurrent

import (
	"context"
	"log"
	"runtime"
)

type GoOption struct {
	FuncName string
	Ctx      context.Context
	Done     chan struct{}
	Extra    map[string]any
}

func (o *GoOption) Go(fn func(o *GoOption)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				log.Printf("recovered from %s panic: %v, stack: %s", o.FuncName, r, string(buf[:n]))
			}
		}()
		fn(o)
		o.Done <- struct{}{}
	}()
}
