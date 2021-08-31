package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	ctx = context.WithValue(ctx, "hello", "hi")
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)

	go runInfiniteLoop(ctx)

	if okay(5) {
		cancel()
	}

	func(x int64) {
		fmt.Println(x)
	}(5)

	for {

	}
}

func runInfiniteLoop(ctx context.Context) {
	defer fmt.Println("hello", ctx.Value("hello"))
	for {
		// do something
		select {
		case <-ctx.Done():
			return
		}
	}
}

func okay(x int64) bool {
	if x&1 == 0 {
		return true
	}
	return false
}
