package main

import "fmt"

const (
	_ = iota
	// kb := 1024
	// kb = 1 << 10 //1이 10번 이동
	// mb = 1024 * kb
	// gb = 1024 * mb
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	//iota 기본
	x := 2
	fmt.Printf("%d\t\t%b", x, x)
	fmt.Println("")
	y := x << 1 // 오른쪽 방향으로 한칸 이동 2->4(이진수)
	fmt.Printf("%d\t\t%b\n", y, y)

	//iota 활용
	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
