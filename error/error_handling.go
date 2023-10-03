package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// go는 try, catch같은 에러처리 방식이 코드를 난해하게 한다고 생각한다. 이유는 일반 파일 열기 실패같은 에러도 예외처리를 하기 때문이다.

// 이 메소드가 붙어있는 타입이라면 이건 반드시 에러이다.
// 사용자 지정 에러 정의 가능
type error interface {
	Error() string
}

func main() {
	// 1
	// var answer1, answer2, answer3 string

	// fmt.Print("Name : ")
	// _, err := fmt.Scan(&answer1)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Print("Fav Food : ")
	// _, err = fmt.Scan(&answer2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Print("Fav Sport : ")
	// _, err = fmt.Scan(&answer3)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(answer1, answer2, answer3)

	//2
	// f, err2 := os.Create("name.txt")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// defer f.Close()

	// r := strings.NewReader("Wassup")

	// io.Copy(f, r)

	//3
	f2, err3 := os.Open("name.txt")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	defer f2.Close()

	bs, err3 := ioutil.ReadAll(f2)
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	fmt.Println(string(bs))
}
