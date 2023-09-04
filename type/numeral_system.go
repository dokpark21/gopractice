package main

import "fmt"

const a = 42 // const는 상수 바꿀 한번 값이 할당되면 바꿀 수 없다
func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)  //2진법
	fmt.Printf("%#X\n", n) //16진법
}
