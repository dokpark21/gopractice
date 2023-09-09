package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	//initialize variable
	// 변수 초기화   조건  : x는 if문 안에서만 사용 가능
	if x := 42; x == 2 {
		fmt.Println("001")
	}

	y := 42
	if y == 40 {
		fmt.Println("value was 40")
	} else if y == 42 {
		fmt.Println("Our value was 42")
	} else {
		fmt.Println("Our value was not 40")
	}

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// && == and, || == or, ! == not,
