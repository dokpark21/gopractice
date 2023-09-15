package main

import "fmt"

// 무엇인가를 defer 하는 것은 파일을 열었을 때 작업이 끝나면 파일을 여는 바로 그 위치에서 파일이 닫히도록 하게 해준다
// 파일을 열고 닫는 부분이 똑같다는 것은 코드의 가독성을 높여준다
// 파일을 연 함수가 어디에 있는지, 파일이 종료되는 위치가 어디인지, 파일을 열고 종료하는 함수가 어디에 있는지에 관례없이 해당 함수에서 몇개의 종료점이 있을 수 있다.
func main() {
	defer foo() // defer 했기 때문에 main func이 종료되는 시점에 실행된다. bar() -> foo() -> end main func
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
