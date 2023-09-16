package main

import "fmt"

// closure : 변수의 범위를 닫고 특정 영역에 담는 것 === 변수의 범위 제한
// 변수의 범위는 최대한 줄여야 한다(에러 발생 줄이기 위해서)

// x의 범위는 함수 전체
var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)

	{
		// y는 이 코드 블럭 안에서만 작동한다
		y := 42
		fmt.Println(y)
	}
	foo()
	fmt.Println(x)

	// 함수가 함수를 반환하기 떄문에 ()()
	a := incrementor()()
	b := incrementor()
	fmt.Println("a : ", a)
	fmt.Println("a : ", a)
	fmt.Println("a : ", a)
	fmt.Println("a : ", a)
	// a는 함수 실행한 결과를 계속 print하는 것이고 b는 함수를 계속 실행한 결과를 print하는 것이기 떄문에 b는 값이 계속 증가한다
	fmt.Println("b : ", b())
	fmt.Println("b : ", b())
	fmt.Println("b : ", b())
	fmt.Println("b : ", b())
}

func foo() {
	x++
}

func incrementor() func() int {
	var x int
	return func() int {
		// x 사용 가능
		x++
		return x
	}
}
