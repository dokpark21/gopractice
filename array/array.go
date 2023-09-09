package main

import "fmt"

func main() {
	var x [5]int //값이 zero로 이루어져 있는 배열 선언
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	//go는 쉬운 프로그래밍, 효울적인 실행, 효율적인 컴파일을 지향하기 떄문에 배열 대신 슬라이스를 사용하라고 명시되어 있습니다.
	//배열은 같은 자료형의 데이터의 연속

	//slice
	// aggregate data type(합성 데이터 타입)은 배열이나 리스트로 불리기도 한다.
	// 역시 같은 데이터의 타입의 데이터를 묶는 것
	// y := type{} // composite literal(합성 리터럴)
	y := []int{4, 5, 6, 7, 8, 42} // 값이 지정되어 잇는 slice 선언
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	//a SLICE allows you to group together VALUES of the same TYPE

	for i, v := range y {
		fmt.Println(i, v)
	}

	//slice slicing
	fmt.Println(y[1:])  // 1~end
	fmt.Println(y[1:3]) // 1~2

	for i := 0; i <= 4; i++ {
		fmt.Println(i, y[i])
	}
}
