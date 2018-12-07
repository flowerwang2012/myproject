package main

import (
	"time"
	"fmt"
	"context"
	"sync"
)

func main() {
	//通过 context.WithValue 来传值
	//key := "key"
	//ctx, cancel := context.WithCancel(context.Background())
	//
	//valueCtx := context.WithValue(ctx, key, "add value")
	//
	//go watch(valueCtx)
	//time.Sleep(10 * time.Second)
	//cancel()
	//
	//time.Sleep(5 * time.Second)
	//通过 context.WithValue 来传值

	//超时取消 context.WithTimeout
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go work(ctx)
	wg.Wait()

	fmt.Println("Finished. I'm going home")
	//超时取消 context.WithTimeout
}

func watch(ctx context.Context) {
	key := "key"
	for {
		select {
		case <-ctx.Done():
			//get value
			fmt.Println(ctx.Value(key), "is cancel")

			return
		default:
			//get value
			fmt.Println(ctx.Value(key), "int goroutine")

			time.Sleep(2 * time.Second)
		}
	}
}

var (
	wg sync.WaitGroup
)

func work(ctx context.Context) error {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing some work ", i)

			// we received the signal of cancelation in this channel
		case <-ctx.Done():
			fmt.Println("Cancel the context ", i)
			return ctx.Err()
		}
	}
	return nil
}

