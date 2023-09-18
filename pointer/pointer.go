package main

import "fmt"

// & : 주소를 확인하는 연산자(앰퍼센트)
// *int, *string 같은 타입은 포인터 타입이다. 기존의 value와는 다르다

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)  // gives you the value stored at an address when you have the address(역참조)
	fmt.Println(*&a) // 역참조(주소를 역참조하면 그 안에 있는 값)

	*b = 43 // 주소에 있는 값을 바꿨기 때문에 a의 값도 같이 바뀐다

	fmt.Println(a)
}
