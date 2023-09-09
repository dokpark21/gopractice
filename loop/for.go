package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// for k := 0; k <= 10; k++ {
	// 	for j := 0; j <= 5; j++ {
	// 		fmt.Println(k + j)
	// 	}
	// }

	//for를 while처럼 사용하기
	// x := 1
	// for x < 10 {
	// 	fmt.Println(x)
	// 	x++
	// }
	// fmt.Println("done")

	// y := 1
	// for {
	// 	if y < 10 {
	// 		break
	// 	}
	// 	fmt.Println(y)
	// 	y++
	// }

	y := 0
	for {
		y++
		if y > 10 {
			break
		}
		if y%2 == 0 {
			continue
		}
		fmt.Println(y)
	}
}
