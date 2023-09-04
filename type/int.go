package main

import "fmt"

// int 8, 16 ,32, 64 float 32, 64 원하는 대로 지정 가능
var x int
var y float64

//int와 uint의 차이점 : u=unsign, int는 음수 가능, uint는 음수 불가능
// ex) int8 : -128~127 uint8: 0~255

func main() {
	x = 42
	y = 42.34532
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}
