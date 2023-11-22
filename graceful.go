package main

import (
	"context"
	"time"
)

func test() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	Consume(ctx)
	time.Sleep(5 * time.Second)
	println("cancel")
	cancel()
	time.Sleep(10 * time.Second)
	println("return")
}

func Consume(ctx context.Context) {
	println("start consume")
	go func() {
		select {
		case <-ctx.Done():
		default:
			time.Sleep(7 * time.Second)
			Handle(ctx)
		}
	}()
}

func Handle(_ context.Context) {
	println("test handle")
}
