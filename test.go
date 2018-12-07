package main

import (
	"time"
	"fmt"
)
type str string
func main() {
	t := time.Now().AddDate(0, 0, -1)
	day := t.Format("2006-01-02")
	fmt.Println(day)
	date, _ := time.ParseInLocation("2006-01-02", day, time.Local)
	url := fmt.Sprintf("%s?date=%d", "www.baidu.com", date.Unix())
	fmt.Println(url)
	s := fmt.Sprintf("[%s] %s", "2017-07-27", "运营数据报表")
	fmt.Println(s)
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println(n)
	fmt.Println("Multiply:", *reply) // Multiply: 50

	st := 1532188800000
	et := 1532584800000
	fmt.Println((et - st) / 1000 / 60 / 60 / 24)

	fmt.Println(st + 1000*60*60*24)

	nums := make([]int, 0)
	nums = append(nums, 1)
	fmt.Println(nums)
	var ss str
	//ss = "abc"
	fmt.Print(ss)

	switch ss {
	case "a":
		ss = "a"
	case "b":
		ss = "b"
	}
	fmt.Print(ss)

	fmt.Println(test2())
}


// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func test2() (n int) {
	n = 1
	return
}