package main

import (
	"fmt"
	"runtime"
	"sync"
)

// runtime package :Go와 관련된 정보와 프로그램이 실행되는 모든 곳에서 실행 중인 CPU의 수를 알 수 있다.

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	// go routine : 1개(main go 루틴)
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// 자체 go routine으로 시작(foo go 루틴)
	// foo는 실행되지 않는다 : go routine이 실행되기 전에 이미 main 함수가 끝나서 프로그램이 종료되기 떄문이다
	// 따라서 기다려달라고 해야 한다(wait group)

	// 1개를 기다리고 있다고 나타낸다
	wg.Add(1)
	go foo()

	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	// go routine : 2개
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// main 함수가 끝나고 foo go 루틴이 실행될 때까지 기다린다.
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo : ", i)
	}
	// 모두 실행되고 나면 종료한다
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar : ", i)
	}
}
