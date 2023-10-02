package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	// receive
	receive(eve, odd, quit)

	fmt.Println("about to exit")
}

func send(eve, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
	quit <- 0
}

func receive(eve, odd, quit <-chan int) {
	for {
		// select 문을 활용해 값을 뺴낸다
		select {
		case v := <-eve:
			fmt.Println("from the even channel:", v)

		case v := <-odd:
			fmt.Println("from the odd channel:", v)

		case v := <-quit:
			fmt.Println("from the quit channel:", v)
			return
		}
	}
}
