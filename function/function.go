package main

import "fmt"

func main() {
	foo()
	bar("Steve")
	s1 := woo("MoneyPenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println("x : ", x)
	fmt.Println("y : ", y)
}

// func (r receiver) identifier (parameters) (return(s)) {...}

func foo() {
	fmt.Println("Hello World!!")
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
