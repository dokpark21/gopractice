package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello golang", 42, true)
	fmt.Println(n) //int 형식 출력 되는 바이트 실행
	fmt.Println(e) //e는 err, 만약 err가 잡힌다면 출력
}

/*
1. golang에서는 버리는 변수가 없어야 한다. 반드시 모두 사용!!
2. 공식docs를 보고 필요한 매개 변수,매개변수의 유형, 들어갈 수 있는 매개 변수의 숫자를 확인
3. {package}.{func} 형식을 사용
*/
