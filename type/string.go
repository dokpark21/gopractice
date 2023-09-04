package main

import "fmt"

func main() {
	s := `"Hello, playground"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%t\n", bs)

	/*
		go의 코드 스킴은 UTF-8
		UTF-8에서 코드 포인트는 1~4바이트
		하나의 코드 포인트는 하나의 문자와 상응한다
		코드 포인트는 16진수로 표현
		rune == int32 == 코드 포인트 == 4바이트(최대)
	*/
	// s에 저장되어 있는 문자열의 UTF-8표현 방식으로 출력
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	//index와 s의 문자를 출력
	for i, v := range s {
		fmt.Println(i, v)
	}
}
