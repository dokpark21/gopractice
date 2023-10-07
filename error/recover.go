package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// recover는 defer된 함수 안에서만 유효하다.
// recover가 panic상태가 아닌 일반 상태에서 호출되면 nil을 반환하며 다른 효과는 없다.
// 실행중인 go routine이 패닉 상태에 빠져 recover를 호출하면 panic에 주어진 값은 저장되고 일반 실행을 재게한다.
