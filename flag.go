package main

import (
	"flag"
	"fmt"
)

// go run flag.go --help 参数说明
// go run flag.go --config ./etc/config.json

// 定义命令行参数
var config = flag.String("config", "config.json", "config path") //返回指针类型
func main() {
	// 解析命令行参数
	flag.Parse()

	fmt.Println(config)
	fmt.Println(*config) //*取值符号
}