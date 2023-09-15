package main

import "fmt"

func main() {
	foo(2, 3, 4, 5, 6, 7)
	bar("Steve")
	s1 := woo("MoneyPenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)

	// 익명 함수
	func(x int) {
		fmt.Println("Anonymous func ran", x)
	}(42)

	// Go에서 함수는 일급객체 : 다른 변수, 다른 타입처럼 똑같이 사용 가능
	f := func(y int) {
		fmt.Println("my first func expression", y)
	}
	f(42)
}

// func (r receiver) identifier (parameters) (return(s)) {...}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Println("Hello World!!")

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding, ", v, "to the total which is now", sum)
	}
	fmt.Println("sum : ", sum)
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("Hello, from woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `,says "Hello"`)
	b := false
	return a, b
}
