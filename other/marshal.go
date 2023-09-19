package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	// json 형식으로 구조체를 marshall하기 위해서는 필드의 첫글자가 반드시 대문자여야 한다
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} else {
		// string으로 형변환해야 필드 명과 값이 프린트 된다
		fmt.Println(string(bs))
	}

}
