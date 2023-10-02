package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	// 여기서는 값을 빼내지 못했기 때문에 ok에서 false가 반환된다.
	v2, ok := <-c
	fmt.Println(v2, ok)
}
