package main

import "fmt"

func main() {
	//6
	// x := 0
	// func() {
	// 	for i := 0; i < 100; i++ {
	// 		x++
	// 	}
	// }()
	// fmt.Println(x)

	//7
	// f := func() {
	// 	fmt.Println("Hello this is Assign func to a variable")
	// }
	// f()
	// fmt.Printf("%T\n", f)

	// var g func()
	// g = func() {
	// 	fmt.Println("helloeeee")
	// }
	// g()

	//8 return func
	// xi := []int{2, 3, 4, 5, 6, 7}
	// f1 := myFunc1(xi...)()
	// fmt.Println(f1)

	//9 callback
	// n := foo(bar, []int{2, 3, 4, 5, 6, 7, 8})
	// fmt.Println(n)

	//10 closure
	f := foo()
	// 실행된 것의 결과를 보여주는 것이 아닌 계쏙 실행을 하고 그 결과를 바로 보여주는 것이기 떄문에 값이 계속 변화한다
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	// 다시 한번
	g := f()
	fmt.Println(g)
	fmt.Println(g)
	fmt.Println(g)
	fmt.Println(g)

}

//8
// 리턴하는 함수가 리턴하는 값의 type까지 꼭 명시해 주어야 하낟
// func myFunc1(xi ...int) func() int {
// 	total := 0
// 	g := func() int {
// 		for _, v := range xi {
// 			total += v
// 		}
// 		return total
// 	}
// 	return g
// }

//9
// func foo(f func(xi []int) int, xi []int) int {
// 	n := f(xi)
// 	return n
// }

// func bar(xi []int) int {
// 	var n int
// 	for _, v := range xi {
// 		n += v
// 	}
// 	return n
// }

// 10
func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func enclosed() {

}
