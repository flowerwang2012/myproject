package main

import (
	"errors"
	"os"
	"os/signal"
	"time"
	"fmt"
)

//这篇通过一个例子，演示使用通道来监控程序的执行时间，生命周期，甚至终止程序等
//
var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrupt = errors.New("执行者被中断")

//一个执行者，可以执行任何任务，但是这些任务是限制完成的，
//该执行者可以通过发送终止信号终止它
type Runner struct {
	tasks     []func(int)      //要执行的任务，同步执行
	complete  chan error       //结果通道，用于通知任务全部完成
	timeout   <-chan time.Time //计时通道，这些任务在多久内完成
	interrupt chan os.Signal   //信号通道，可以控制强制终止的信号
}

func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal),
	}
}

//将需要执行的任务，添加到Runner里
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//执行任务，执行的过程中接收到中断信号时，返回中断错误
//如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r *Runner) run() error {
	go r.isInterrupt()
	for id, task := range r.tasks {
		task(id)
	}
	return nil
}

//检查是否接收到了中断信号
func (r *Runner) isInterrupt() {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		r.complete <- ErrInterrupt
	}
}

// 开始执行所有任务，并且监视通道事件
// 两种情况，要么任务完成；要么到时间了，任务执行超时。
// 任务完成又分两种情况，一种是没有执行完，但是收到了中断信号，中断了，这时返回中断错误；一种是顺利执行完成，这时返回nil。
func (r *Runner) Start() error {
	//监听接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()
	//多路复用
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func main() {
	fmt.Println("...开始执行任务...")

	timeout := 3 * time.Second
	r := New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			fmt.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			fmt.Println(err)
			os.Exit(2)
		}
	}
	fmt.Println("...任务执行结束...")
}

func createTask() func(int) {
	return func(id int) {
		fmt.Printf("正在执行任务%d\n", id)
		time.Sleep(time.Duration(id) * 2 * time.Second)
		fmt.Printf("任务%d执行完毕\n", id)
	}
}
