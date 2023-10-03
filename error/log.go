package main

import (
	"fmt"
	"log"
	"os"
)

// 에러 출력
func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// 만약 log가 찍히면 f(log.txt)에 쓰고 저장
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		// log.Println("error happened : ", err)

		// log를 찍고 exit 함수 호출(바로 프로그램 종료)
		log.Fatalln(err)
	}
	fmt.Println("Is this print or not print?? Check this out!") // 위에서 log.Fatalln이 exit를 호출해서 이 코드는 실행되지 않고 종료 된다.
	defer f2.Close()
}
