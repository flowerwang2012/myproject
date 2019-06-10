package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	//test1 "myproject/test"
)

func main() {
	fmt.Println(os.Args) //[/var/folders/16/2v226hy904v81jvwrvrvs8940000gn/T/go-build077640365/b001/exe/glog -config etc/config.json -key value]
	fmt.Println(os.Args[1:]) //[-config etc/config.json -key value]
	configPtr := flag.String("config", "1", "") //flag.string(...)会将 指定名字的flag对象 缓存在flag中
	valuePtr := flag.String("key", "2", "")
	flag.Parse()
	fmt.Println(*configPtr) //etc/config.json
	fmt.Println(*valuePtr) //value

	//test1.LogDirTest()
	//flag.Lookup("log_dir").Value.Set("./") //如果有定义的对象，flag.string(...)，那Lookup的对象就不为空
	//fmt.Println(flag.Lookup("log_dir").Value.String())

	// glog package定义了log_dir的flag结构体，这里直接给结构体对象的value赋值
	flag.Lookup("log_dir").Value.Set("./") //指定glog执行日志写入的目录
	flag.Lookup("v").Value.Set("2") //指定glog执行日志写入的等级

	glog.Info("this is a info log")
	// 方法调用后，产生两个文件，但是没有写入数据：
	// glog.INFO 文件
	// glog.wanghongfadeMacBook-Pro.wanghongfa.log.INFO.20190510-195704.1565 文件
	// 需要显示地调用 glog.Flush()
	defer glog.Flush()
	// 因为数据会暂存在内存的buffer中。只有显示的调用 glog.Flush(), 数据才会真正被写入文件。
	// glog package 的 init 函数中启动了一个 goroutine (go logging.flushDaemon())用来周期性的调用 glog.Flush() 来保证数据被写入文件, 默认的 Flush 周期为30 秒。

	// 这两个方法调用要生效必须指定flag v的值，当v的值大于等于V()里面的参数才执行后面的Info，如果不加-v参数，默认等级为0。
	glog.V(1).Info("level 1")
	glog.V(2).Info("level 2")
}