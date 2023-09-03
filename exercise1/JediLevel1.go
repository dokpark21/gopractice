package main

import "fmt"

var a int
var b string
var c bool

// package level scope(main package)
var one int = 42
var two string = "James Bond"
var three bool = true

type hotdog int

var v hotdog = 21

var r int

func main() {
	//1
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)

	//2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//a b c is zero value

	//3
	s := fmt.Sprintf("%v\t%v\t%v", one, two, three)
	fmt.Println(s)

	//4
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	v = 42
	fmt.Println(v)
	fmt.Printf("%T\n", v)

	//5
	r = int(v)
	fmt.Println(r)
	fmt.Printf("%T\n", r)
}
