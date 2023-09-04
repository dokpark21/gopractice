package main

import "fmt"

const a = 42 // const는 상수 바꿀 한번 값이 할당되면 바꿀 수 없다, 상수는 타입이 없다. 컴파일러에게 유연성을 준다. 다른 무언가에 지정해 줄 수 있기 때문에 유연성이 생긴다

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
