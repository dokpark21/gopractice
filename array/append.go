package main

import "fmt"

func main() {

	x := []int{4, 5, 7, 8, 9}
	fmt.Println(x)
	// 배열이나 슬라이스는 무저건 앞에 와야한다
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// 값 지우기
	// x 0~1, 4~끝까지
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
