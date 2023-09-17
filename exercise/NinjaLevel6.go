package main

import (
	"fmt"
	"math"
)

// type person struct {
// 	first string
// 	last  string
// 	age   int
// }

// func (p person) speak() {
// 	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old")
// }

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	// return 해주는 값이 있기 떄문에 float64를 지정해줘야 한다
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}

func main() {
	// s1 := foo()
	// s2, s3 := bar()
	// fmt.Println(s1, s2, s3)

	// xi := []int{2, 3, 4, 5, 65}
	// sum := foo(xi...)
	// fmt.Println(sum)

	// sum2 := bar(xi)
	// fmt.Println(sum2)

	// foo()

	// p1 := person{
	// 	first: "Park",
	// 	last:  "steve",
	// 	age:   25,
	// }
	// p1.speak()

	circ := circle{
		radius: 2.5,
	}

	squa := square{
		length: 32.12,
	}

	info(circ)
	info(squa)
}

//1
// func foo() int {
// 	a := 42
// 	return a
// }

// func bar() (int, string) {
// 	b := 43
// 	s := "string"
// 	return b, s
// }

// 2
// func foo(x ...int) int {
// 	total := 0
// 	for _, v := range x {
// 		total += v
// 	}
// 	return total
// }

// func bar(xx []int) int {
// 	total := 0
// 	for _, v := range xx {
// 		total += v
// 	}
// 	return total
// }

// 3
// func foo() {
// 	defer fmt.Println("This is first line of func foo but run after func is over")
// 	fmt.Println("This is second line")
// 	fmt.Println("This is third line of func foo!!")
// 	fmt.Println("This is fourth line!")
// }

//4
