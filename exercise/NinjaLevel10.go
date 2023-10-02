package main

import (
	"fmt"
)

func main() {
	//1
	// c := make(chan int)

	// go func() {
	// 	c <- 42
	// }()

	// fmt.Println(<-c)

	//2
	// cs := make(chan int)
	// go func() {
	// 	cs <- 42
	// }()
	// fmt.Println(<-cs)

	//3
	// c := gen()
	// receive(c)

	// fmt.Println("about to exist")

	//4
	// q := make(chan int)
	// c := gen(q)

	// receive(c, q)

	// fmt.Println("about to exist")

	//5
	// c := make(chan int)
	// go func() {
	// 	c <- 42
	// 	c <- 43
	// 	close(c)
	// }()

	// v, ok := <-c
	// fmt.Println(v, ok)

	// v, ok = <-c
	// fmt.Println(v, ok)

	// v, ok = <-c
	// fmt.Println(v, ok)

	//6,7
	c := make(chan int)

	for g := 0; g < 10; g++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("about to exit")

}

//3
// func gen() <-chan int {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 		close(c)
// 	}()
// 	return c
// }

// func receive(c <-chan int) {
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }

// 4
// func gen(q chan<- int) <-chan int {
// 	c := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			c <- i
// 		}
// 		q <- 1
// 		close(c)
// 	}()

// 	return c
// }

// func receive(c, q <-chan int) {
// 	for {
// 		select {
// 		case v := <-c:
// 			fmt.Println(v)
// 		case <-q:
// 			return
// 		}
// 	}
// }
