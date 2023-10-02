package main

// fanin : 여러개의 다른 채널의 값들을 하나의 fanin 채널로 모아준다.
// 다중 패널을 하나의 채널로 팬인
import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// even과 odd는 값을 받기 때문에 송신 전용 채널로 특정해 준다.
func send(even, odd chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// even과 odd에서 값을 빼서 fanin에 넣어줘야 하기 때문에 각각 수신과 송신 채널로 특정해 준다.
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
