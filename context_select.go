package main

import (
	"context"
	"fmt"
	"time"
)
// 启动了3个监控goroutine进行不断的监控，每一个都使用了Context进行跟踪，
// 当我们使用cancel函数通知取消时，这3个goroutine都会被结束。
// 这就是Context的控制能力，它就像一个控制器一样，按下开关后，
// 所有基于这个Context或者衍生的子Context都会收到通知，这时就可以进行清理操作了，最终释放goroutine，
// 这就优雅的解决了goroutine启动后不可控的问题
func main() {
	// context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点，这个空的Context不能被取消、没有值、也没有过期时间
	// context.TODO() 与Background()的返回值一样，从字面理解就是你不确定使用什么context，就用TODO()，仅此而已
	// 然后我们使用context.WithCancel(parent)函数，创建一个可取消的子Context，
	// 然后当作参数传给goroutine使用，这样就可以使用这个子Context跟踪这个goroutine。
	ctx, cancel := context.WithCancel(context.Background())
	go watcher2(ctx, "监控1号")
	go watcher2(ctx, "监控2号")
	go watcher2(ctx, "监控3号")
	time.Sleep(3 * time.Second)
	fmt.Println("可以了，停止所有监控")
	cancel()
	time.Sleep(3 * time.Second)
	fmt.Println("所有监控已经停止")
}

func watcher2(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("停止%s...\n", name)
			return
		default:
			fmt.Printf("%s工作中...\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}
