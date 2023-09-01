package main

import "fmt"

func main() {
	//변수를 처음 선언할 때는 :=사용
	x := 42
	fmt.Println(x)
	//변수를 재선언할 때는 = 사용
	x = 35
	fmt.Println(x)
	//변수는 식으로 지정 가능
	y := 100 + 24
	//여기서 100과 24는 operand라고 부른다. operand = 수행해야 할 동작을 표현하는 작은 요소
	//구문 = 프로그램을 구성하는 요소들 보통은 한 줄을 지칭 만약 한줄에 두개가 존재한다면 세미콜론으로 구분
	fmt.Println(y)
}

//자세한 연산자, keywords등은 goalng spec에서 확인 가능
