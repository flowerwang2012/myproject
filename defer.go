package main

import "fmt"

//https://my.oschina.net/henrylee2cn/blog/505535
func main() {
	fmt.Println(Parse())
}

func Parse() (err error) {
	defer func() { //延迟执行语句，函数结束时执行，实现错误捕捉和恢复
		//recover，宕机恢复
		switch p := recover(); p {
		case nil:
		case "what":
			err = fmt.Errorf("internal error")
		default:
			panic(p)
		}
	}()
	defer fmt.Println("延迟执行栈，这个先出栈")
	panic("what") //宕机
	fmt.Println(1)
	return err
}
