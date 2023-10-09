package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		// r은 패닉에 주어진 값을 저장하는 변수이다.
		// r이 nil이 아니리면 패닉에 빠져 값을 저장한 것이기 떄문이다.
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
	// defer은 후입 선출이다.
	// i가 3이 넘지 않으면 함수가 패닉 상태에 빠져 종료되지 않고 계속해서 g(i+1)을 호출하고 스택에는 defer함수들이 쌓이게 된다.
	// 만약 i가 3을 넘어서 함수가 패닉 상태에 빠진다면 스택에서 쌓인 defer함수들을 하나씩 실행시킨다.(후입 선출)
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// recover는 defer된 함수 안에서만 유효하다.
// recover가 panic상태가 아닌 일반 상태에서 호출되면 nil을 반환하며 다른 효과는 없다.
// 실행중인 go routine이 패닉 상태에 빠져 recover를 호출하면 panic에 주어진 값은 저장되고 일반 실행을 재게한다.
