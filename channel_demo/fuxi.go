package main

import (
	"fmt"
	"math/rand"
	"time"
)

//并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）。
//并行：同一时刻执行多个任务（你和你朋友都在用微信和女朋友聊天）

// goroutine 启动
// 将要并发执行的任务包装成一个函数，调用函数的时候前面加上一个go关键字，就能够开启一个goroutine
// 去执行该函数的任务
// goroutine对应的函数执行完，该goroutine就结束了
// 程序启动的时候会自动创建一个goroutine去执行main函数
// main函数结束了，那么程序就结束了，由该程序启动的所有开启的goroutine也都结束了

// goroutine 的本质
// goroutine的调度模型： GMP
// 把m个goroutine分配给n个操作系统线程

// goroutine 与操作系统线程（os线程）的区别
// goroutine 是用户态的线程， 比内核态的线程更轻量级一点，初始值只占用2KB的空间
// 可以轻松开启数十万个goroutine的也不会崩内存

// runtime.GOMAXPROCS
// GO1.5之后默认就是操作系统cpu的逻辑核心数，默认跑满cpu
// runtime.GOMAXPROCS(1) : 只占用一个核 （收集日志，监控系统最好设置）

// wook pool
// 开启一定数量的goroutine去干活

// sync.WaitGroup
// var wg sync.WaitGroup
// wg.Add(n) 计数器+1
// wg.Done() 计数器-1
// wg.Wait() 等

// channel
// 通过channel实现多个goroutine之间的通信，从而达到协同工作的目的
// CSP: 通过通信来共享内存
// channel 声明/初始化  var ch chan 元素类型  / ch = make(chan 元素类型 ， [缓冲区大小])
// 发送 ch <-
// 接收 x := <- ch
// 关闭 close(ch) , 非必须
// 带缓冲区的通道和无缓冲区的通道

func numRand(ch chan<- int64) {
	for {
		n := rand.Int63n(10)
		ch <- n
		time.Sleep(time.Second * 5)
	}
}
func main() {

	ch := make(chan int64, 1)

	go numRand(ch)
	for {
		n, ok := <-ch

		if !ok {
			break
		}

		fmt.Println(n, ok)
		time.Sleep(time.Second)
	}

}
