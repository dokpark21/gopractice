package main

import "fmt"

// var는 함수 외부에서 선언가능
// 프로그램 전체에서 사용가능
// 이렇게 전역변수로 저장하게 되면 겹치는 곳이 생겨서 에러 발생 가능성이 올라간다
var y = 43

// 이렇게 지정하게 되면 zero value로 저장된다
var z int

func main() {

	// DECLARE avariable and ASSIGN a VALUE
	// 단축 선언 연산자는 함수 내부에서만 선언가능 및 사용 가능, 변수의 범위를 최대한 제한가능하다. 최대한 단축 선언 연산자를 사용하는 것이 좋다.
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println(y)
}
