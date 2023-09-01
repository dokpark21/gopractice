package main

import "fmt"

/*
go언어에서는 타입이 굉장히 중요하다
go는 정적 언어이기 때문에 변수 형을 변경할 수 없다
*/

var y = 42

var a string = `Steve said, "I'm a most valueable basketball player in Suwon"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
