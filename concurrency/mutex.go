package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs : ", runtime.NumCPU())
	fmt.Println("Go routines : ", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// 코드를 잠궈서 counter 변수에 아무도 접근하지 못하게한다. == 동시에 여러개의 루틴이 접근 불가능
			// v++ -> counter = v로 넘어가는 과정에서 만약 다른 go routine이  v:= counter를 실행하게 되면서 값이 늘어난 것이 다시 원상태로 돌아온 것이다.
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			// 동작이 끝나면 잠금을 해제한다.
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Go routines : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines : ", runtime.NumGoroutine())
	fmt.Println("Counter : ", counter)
}
