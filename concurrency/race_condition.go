package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 아래의 코드를 실행시켜 보면 go routine의 개수가 오르락 내리락하고 Counter도 100이 아닌 숫자가 계속 바뀔 것이다.
// 이런 이유는 경쟁상태에 있어서이다

func main() {
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("Go routines : ", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// 1초 기다린다
			// time.Sleep(time.Second)
			// 다른 go routine에게 양보(yield)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Go routines : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines : ", runtime.NumGoroutine())
	fmt.Println("Counter : ", counter)
}


// go run -race로 실행해보면 1개의 레이스가 있다. 레이스가 실행되고 있는 이유는 다수의 Go routine이 있기 때문이다.
