package main

import "fmt"

func main() {
	fmt.Println("Hello World!!")
	//go에서는 컴파일 시 자동으로 세미콜론을 삽입해준다. 따라 삽입할 필요가 없다.

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println((i))
		}
	}
}

func foo() {
	fmt.Println("This is foo func!!!")
}

/*
workflow
1. package main
2. func main
3. print hello world
4. func foo
5. print this is foo func
6. print i if i%2==0
*/
