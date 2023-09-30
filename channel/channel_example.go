package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go foo(c)
	// receive
	// 이렇게 해야 bar 함수는 자신이 실행될 때 까지 기다렸다 출력, 그리고 main 함수 종료
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
