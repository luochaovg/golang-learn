package main

import (
	"fmt"
	"sync"
)

//交替打印数字和字⺟
//问题描述
//使⽤两个  goroutine 交替打印序列，⼀个  goroutine 打印数字， 另外⼀ 个  goroutine 打印字⺟， 最终效果如下：

//解题思路
//问题很简单，使⽤ channel 来控制打印的进度。使⽤两个 channel ，来分别控制数字和
//字⺟的打印序列， 数字打印完成后通过 channel 通知字⺟打印, 字⺟打印完成后通知数
//字打印，然后周⽽复始的⼯作。
//

var wg sync.WaitGroup

func main() {
	letter, number := make(chan bool), make(chan bool)

	i := 0
	go func() {
		for {
			select {
			case <-number:
				i++
				fmt.Println(i)

				i++
				fmt.Println(i)

				letter <- true
				break
			default:
				break

			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

		for {
			select {
			case <-letter:
				strbyts := []byte(str)

				if i-2 >= len(strbyts) {
					return
				}

				one := strbyts[i-2]
				two := strbyts[i-1]

				fmt.Println(string(one))
				fmt.Println(string(two))

				number <- true
				break

			default:
				break

			}
		}

	}()

	number <- true
	wg.Wait()

}
