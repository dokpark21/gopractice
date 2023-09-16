package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	// bar는 int를 반환하는 함수를 반환
	s2 := bar()
	// 따라서 한번 더 인수에 함수 처럼 할당해야 return 되는 값을 사용할 수 있다.
	x := s2()
	fmt.Println(x)
	fmt.Println(s2())
}

// returning string
func foo() string {
	s := "Hello World"
	return s
}

// int를 반환하는 func를 return
func bar() func() int {
	return func() int {
		return 451
	}
}
