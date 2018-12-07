package main

import "fmt"

//https://my.oschina.net/henrylee2cn/blog/505535
func main() {
	fmt.Println(Parse())
}

func Parse()(err error){
	defer func() {
		//选择性的recover
		switch p := recover(); p {
		case nil:
		case "what":
			err = fmt.Errorf("internal error")
			//fmt.Println(err)
		default:
			panic(p)

		}
	}()
	panic("what")
	return err
}