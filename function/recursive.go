package main

import "fmt"

// Recursive : 함수가 자기 자신을 호출한다 == loop
// loop도 동일한 동작을 하고 메모리에 안좋은 영향을 끼친다, 좀 더 복잡함
func main() {
	n := factorial(4)
	fmt.Println(n)
}

func factorial(n int) int {
	// 만약 1이면 다시 전에 미뤄두었던 동작들을 실행
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
