package main

import "fmt"

// boolean = true or false
var x bool //zero value == false

func main() {
	fmt.Println(x)
	x = true
	a := 7
	b := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(a == b)
	fmt.Println(a <= b)
}
