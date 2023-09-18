package main

import "fmt"

// & : 주소를 확인하는 연산자(앰퍼센트)
// *int, *string 같은 타입은 포인터 타입이다. 기존의 value와는 다르다

func main() {
	// a := 42
	// fmt.Println(a)
	// fmt.Println(&a) // & gives you the address

	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", &a)

	// b := &a
	// fmt.Println(b)
	// fmt.Println(*b)  // gives you the value stored at an address when you have the address(역참조)
	// fmt.Println(*&a) // 역참조(주소를 역참조하면 그 안에 있는 값)

	// *b = 43 // 주소에 있는 값을 바꿨기 때문에 a의 값도 같이 바뀐다

	// fmt.Println(a)

	x := 0
	fmt.Println(x)
	fmt.Println(&x)
	foo(&x)
	// foo 에서 주소를 참조해 값을 변경했기 때문에 값이 바뀐다(주소는 변하지 않는다)
	fmt.Println(&x)
	fmt.Println(x)
}

/*
포인터는 해당 데이터가 저장된 주소만 전달하면 된다.
1. 특정 위치에 무엇인가 변경해야 할 때(mutate)
*/

func foo(x *int) {
	fmt.Println(*x)
	// 여기서 값을 바꾼다고 main에서 값이 바뀌지 않는다.
	// 바꿀려면 주소를 참조해 값을 바꿔야 한다.
	*x = 43
	fmt.Println(x)
}

// 타입에 연결하는 메소드를 메소드 세트라고 한다
// 포인터 receiver은 포인터에만 작동한다
