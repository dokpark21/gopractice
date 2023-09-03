package main

import "fmt"

var y = 42

func main() {
	n, e := fmt.Println("Hello golang", 42, true)
	fmt.Println(n) //int 형식 출력 되는 바이트 실행
	fmt.Println(e) //e는 err, 만약 err가 잡힌다면 출력

	fmt.Println(y)
	fmt.Printf("%T\n", y)  //y type
	fmt.Printf("%b\n", y)  //y binary(2진수)
	fmt.Printf("%x\n", y)  //y 16진수
	fmt.Printf("%#x\n", y) //y 0x16진수
	//Println은 한줄 띄어서 출력, Printf는 포맷 프린트

	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y) //Sprint는 문자열 프린트
	fmt.Println(s)
}

/*
1. golang에서는 버리는 변수가 없어야 한다. 반드시 모두 사용!!
2. 공식docs를 보고 필요한 매개 변수,매개변수의 유형, 들어갈 수 있는 매개 변수의 숫자를 확인
3. {package}.{func} 형식을 사용
*/
