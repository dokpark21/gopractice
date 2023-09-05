package main

import "fmt"

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
	fmt.Println("here's a statement")
	fmt.Println("something else")
}
