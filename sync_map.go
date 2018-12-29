package main

import (
	"sync"
	"time"
	"fmt"
)

func main() {
	//并发的map读写报错：
	//fatal error: concurrent map read and map write
	//需要并发读写时，一般做法时加锁，但这样性能并不高，go提供了一种效率较高的并发安全的sync.Map
	//m := make(map[int]int)
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		m[i] = i
	//	}
	//}()
	//go func() {
	//	for i := 0; i < 1000; i++ {
	//		_ = m[i]
	//	}
	//}()
	//time.Sleep(2 * time.Second)

	var m sync.Map
	go func() {
		for i := 0; i < 1000; i++ {
			m.Store(i,i)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			m.Load(i)
		}
	}()
	time.Sleep(2 * time.Second)
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}