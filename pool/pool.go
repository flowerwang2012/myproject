package main

import (
	"fmt"
	"sync"
	"io"
	"sync/atomic"
	"errors"
	"time"
	"math/rand"
)

// 资源池，不限制资源类型
type Pool struct {
	m       sync.Mutex
	rw      sync.RWMutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

var ErrPoolClosed = errors.New("资源池已经被关闭。")

// 初始化资源池
func New(f func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了。")
	}
	p := &Pool{
		res:     make(chan io.Closer, size),
		factory: f,
	}
	for i := 0; i < size; i++ {
		r, err := p.factory()
		if err != nil {
			return p, err
		}
		p.res <- r
	}
	fmt.Println("初始化资源池")
	return p, nil
}

// 获取资源 读操作 没有资源阻塞等待其他goroutine释放资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		fmt.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	}
}

// 释放资源 写操作
// 因为Close和Release这两个方法是互斥的，Close方法里对closed标志的修改，Release方法可以感知到，所以就直接return了，不会执行下面的select代码了，也就不会往一个已经关闭的通道里发送资源了。
func (p *Pool) Release(r io.Closer) {
	//保证该操作和Close方法的操作是安全的，当goroutine在修改closed标志时，其他goroutine都不能进行读操作
	//但是会有一个问题，就是在这里只能有一个goroutine在执行读操作，显然是不理想的
	//理想状态是一个goroutine修改closed标志时，其他goroutine都不能进行读操作，修改完毕后，其他goroutine能并发的读
	//所以，这个场景我们需要用到读写锁，可以参考rwmutex.go
	//p.m.Lock()
	//defer p.m.Unlock()
	p.rw.RLock()
	defer p.rw.RUnlock()
	//资源池都关闭了，就省这一个没有释放的资源了，释放即可
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		fmt.Println("资源释放到池子里了")
	default:
		fmt.Println("资源池满了，释放这个资源吧")
		r.Close()
	}
}

// 关闭资源
func (p *Pool) Close() {
	//p.m.Lock()
	//defer p.m.Unlock()
	p.rw.Lock()
	defer p.rw.Unlock()

	if p.closed {
		return
	}
	p.closed = true
	//关闭通道，不让写入了
	close(p.res)
	//关闭通道里的资源
	for r := range p.res {
		r.Close()
	}
}

const (
	//模拟的最大goroutine
	maxGoroutine = 5
	//资源池的大小
	poolRes = 2
)

func main() {
	//等待任务完成
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
	p, err := New(createConnection, poolRes)
	if err != nil {
		fmt.Println(err)
		return
	}
	//模拟好几个goroutine同时使用资源池查询数据
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	fmt.Println("开始关闭资源池")
	p.Close()
}

//模拟数据库查询
func dbQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pool.Release(conn)
	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("第%d个查询，使用的是ID为%d的数据库连接\n", query, conn.(*dbConnection).ID)
}

// 实现io.Closer接口的资源
type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	fmt.Println("关闭对象资源")
	return nil
}

var idCounter int32

// 创建对象资源
func createConnection() (io.Closer, error) {
	//并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}
