package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
	"runtime/trace"
)

var wg sync.WaitGroup

func init()  {
	rand.Seed(time.Now().UnixNano())
}

func main()  {
	trace.Start(os.Stderr)
	defer  trace.Stop()

	court := make(chan int)

	wg.Add(2)

	go player("Luochao", court)
	go player("xiaoxue", court)

	court <- 1

	wg.Wait()


}

func player(name string, court chan int)  {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Printf("Player %s Won \n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed \n", name)

			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d \n", name, ball)
		ball++

		court <- ball

	}
}
