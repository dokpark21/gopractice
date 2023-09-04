package main

import "fmt"

const (
	a = iota //iota는 미리 지정된 특별한 상수이다. a를 iota로 지정해 놓으면 b,c는 iota를 지정해주지 않아도 알아서 증가한다(const로 묶인 것 안에서만 작동)
	b = iota
	c = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
