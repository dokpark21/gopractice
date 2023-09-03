package main

import "fmt"

/*
go언어에서는 타입이 굉장히 중요하다
go는 정적 언어이기 때문에 변수 형을 변경할 수 없다
*/

var y = 42

// 이렇게 아무 것도 지정하지 않는다면 0 = zero value
var z int

var a string = `Steve said, "I'm a most valueable basketball player in Suwon"`

// 나만의 타입 생성
var x int

// hotdog라는 나만의 type 생성
type hotdog int

// b의 type은 hotdog
var b hotdog

func main() {
	fmt.Println(y)
	//타입 확인
	fmt.Printf("%T\n", y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	x = 42
	b = 43
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(b)        //43
	fmt.Printf("%T\n", b) //main.hotdog === main package에 hotdog type

	//type 변환

	x = int(b) // hotdog -> int로 변환해 x에 b의 값을 할당한다
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
