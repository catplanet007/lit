package main

import (
	"context"
	"log"
	"time"

	lcc "github.com/catplanet007/lit/lconcurrent"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	opt := lcc.GoOption{
		FuncName: "my-func",
		Ctx:      ctx,
		Done:     make(chan struct{}),
	}
	opt.Go(func(o *lcc.GoOption) {
		log.Printf("hello %s", o.FuncName)
		select {
		case <-o.Ctx.Done():
			log.Printf("ctx canceld: %s", o.Ctx.Err().Error())
		case <-time.After(time.Second * 2):
		}
		panic("my-error")
	})
	<-opt.Done
}
