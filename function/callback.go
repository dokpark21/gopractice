package main

import "fmt"

// callback = 인수로 함수를 전달

func main() {
	ii := []int{2, 3, 4, 5, 6, 7, 8}
	s := sum(ii...)
	fmt.Println("all number :", s)
	// 함수를 인자로 전달할때는 함수명만 사용하면 된다
	s2 := even(sum, ii...)
	fmt.Println("even number :", s2)
}

func sum(xi ...int) int {
	total := 0
	fmt.Printf("%T\n", xi)
	for _, v := range xi {
		total += v
	}

	return total
}

// 함수를 인자로 취하는 콜백
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	sum := f(yi...)

	return sum
}
