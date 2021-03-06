package main

import (
	"runtime"
	"fmt"
	"sync"
)
/*
在操作系统中，有两个重要的概念：一个是进程、一个是线程。
进程是程序的工作空间，包含了运行这个程序所需的各种资源。线程是一个执行的空间，比如要下载一个文件，访问一次网络等等。
一个进程在启动的时候，会创建一个主线程，这个主线程结束的时候，程序进程也就终止了，所以一个进程至少有一个线程，这也是我们在main函数里，使用goroutine的时候，要让主线程等待的原因，因为主线程结束了，程序就终止了，那么就有可能会看不到goroutine的输出。
go语言中并发指的是让某个函数独立于其他函数运行的能力，一个goroutine就是一个独立的工作单元，Go的runtime（运行时）会在逻辑处理器上调度这些goroutine来运行，一个逻辑处理器绑定一个操作系统线程，所以说goroutine不是线程，它是一个协程，也是这个原因，它是由Go语言运行时本身的算法实现的。
当我们创建一个goroutine的后，会先存放在全局运行队列中，等待Go运行时的调度器进行调度，把他们分配给其中的一个逻辑处理器，并放到这个逻辑处理器对应的本地运行队列中，最终等着被逻辑处理器执行即可。
这一套管理、调度、执行goroutine的方式称之为Go的并发。并发可以同时做很多事情，比如有个goroutine执行了一半，就被暂停执行其他goroutine去了，这是Go控制管理的。
所以并发的概念和并行不一样，并行指的是在不同的物理处理器上同时执行不同的代码片段，并行可以同时做很多事情，而并发是同时管理很多事情，因为操作系统和硬件的总资源比较少，所以并发的效果要比并行好的多，使用较少的资源做更多的事情，也是Go语言提倡的。
*/
func main() {
	fmt.Println(runtime.NumCPU()) //cpu处理器数量
	fmt.Println(runtime.NumGoroutine()) //运行时的goroutine数量
	runtime.GOMAXPROCS(1)//设置逻辑处理器
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i:=1;i<100000;i++ {
			fmt.Println("A:",i)
		}
	}()
	go func() {
		defer wg.Done()
		for i:=1;i<100000;i++ {
			fmt.Println("B:",i)
		}
	}()
	wg.Wait()
	/*
	我们运行这个程序，会发现A和B前缀会交叉出现，并且每次运行的结果可能不一样，这就是Go调度器调度的结果。
	默认情况下，Go默认是给每个可用的物理处理器都分配一个逻辑处理器，因为我的电脑是4核的，所以上面的例子默认创建了4个逻辑处理器，所以这个例子中同时也有并行的调度，如果我们强制只使用一个逻辑处理器，我们再看看结果。
	并发的效果
	B: 99997
	B: 99998
	B: 99999
	A: 56343
	A: 56344
	A: 56345
	 */
}
/*
对于逻辑处理器的个数，不是越多越好，要根据电脑的实际物理核数，如果不是多核的，设置再多的逻辑处理器个数也没用，如果需要设置的话，一般我们采用如下代码设置。
runtime.GOMAXPROCS(runtime.NumCPU())
所以对于并发来说，就是Go语言本身自己实现的调度，对于并行来说，是和运行的电脑的物理处理器的核数有关的，多核就可以并行并发，单核只能并发了。
 */