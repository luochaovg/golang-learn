package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	rand.Seed(time.Now().UnixNano()) // 保证每次执行都有点不一样

	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(100) // 0 <= x < 10
		fmt.Println(r1, 0-r2)
	}

}

func f1(i int) {
	defer wg.Done()

	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

// struct 值类型
var wg sync.WaitGroup

// waitGroup

// 程序启动之后会创建一个main goroutine 去执行
func main() {
	//f()

	// 如何知道这10 个goroutine都结束了
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	wg.Wait() // 等待wg的计数器减为0

	fmt.Println("main over")
}

// https://www.liwenzhou.com/posts/Go/14_concurrence/
// GMP goroutine 调度模型
//goroutine与线程
//可增长的栈
//OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,
// 一个goroutine的栈在其生命周期开始时只有很小的栈（典型情况下2KB），
// goroutine的栈不是固定的，他可以按需增大和缩小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这么大。
// 所以在Go语言中一次创建十万左右的goroutine也是可以的。
//
//goroutine调度
//GPM是Go语言运行时（runtime）层面的实现，是go语言自己实现的一套调度系统。区别于操作系统调度OS线程。
//
//G很好理解，就是个goroutine的，里面除了存放本goroutine信息外 还有与所在P的绑定等信息。
//P管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界），
//P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等）当自己的队列消费完了就去全局队列里取，
// 如果全局队列里也消费完了会去其他P的队列里抢任务。
//M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个groutine最终是要放到M上执行的；
//P与M一般也是一一对应的。他们关系是： P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，
// runtime会新建一个M，阻塞G所在的P会把其他的G 挂载在新建的M上。当旧的G阻塞完成或者认为其已经死掉时 回收旧的M。
//
//P的个数是通过runtime.GOMAXPROCS设定（最大256），Go1.5版本之后默认为物理线程数。 在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。
//
//单从线程调度讲，Go语言相比起其他语言的优势在于OS线程是由OS内核来调度的，
// goroutine则是由Go运行时（runtime）自己的调度器调度的，
// 这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。
// 其一大特点是goroutine的调度是在用户态下完成的， 不涉及内核态与用户态之间的频繁切换，
// 包括内存的分配与释放，都是在用户态维护着一块大的内存池， 不直接调用系统的malloc函数
// （除非内存池需要改变），成本比调度OS线程低很多。 另一方面充分利用了多核的硬件资源，
// 近似的把若干goroutine均分在物理线程上， 再加上本身goroutine的超轻量，以上种种保证了go调度方面的性能。
