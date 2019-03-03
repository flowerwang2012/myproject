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
	defer cancel() //cancel函数在这里的意义不大，所以用defer延迟函数去执行，可以用占位符_取代cancel返回值，也能正常执行

	fmt.Println("Hey, I'm going to do some work")

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := work(ctx); err != nil {
			fmt.Println(err.Error())
		}
	}()
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

