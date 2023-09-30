package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)
	//receive
	for v := range c {
		fmt.Println(v)
	}
}

func foo(c chan<- int) {
	// 채널이 닫히기 전까지 채널에서 값을 빼낸다.
	// 닫힌 채널에 더 이상 값이 없으면 채널을 빠져 나온다.
	for i := 0; i < 100; i++ {
		c <- i
	}
	// 채널을 닫아준다
	// 채널을 닫지 않으면 값이 더 올 때 까지 기다리기 떄문에 교착 상태에 빠진다.
	close(c)
}
