package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// fanout : 하나의 작업을 여러개의 go routine으로 팬아웃(분할하여 동시 실행)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		// 100개의 작업을 하나씩 waitgroup을 부여(하나씩 진행하는 것이 아닌 100개로 쪼개서 여러개 동시 진행)
		wg.Add(1)

		go func(v2 int) {
			// 랜덤한 값 부여
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(500)
}
