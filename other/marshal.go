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

// marshal : go type struct => JSON
// unmarshal : JSON => go type struct(JSON으로 인코딩된 data를 분석하고 v가 가르키는 곳에 저장)

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

	// json 형식으로 구조체를 marshal하기 위해서는 필드의 첫글자가 반드시 대문자여야 한다
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	} else {
		// string으로 형변환해야 필드 명과 값이 프린트 된다
		fmt.Println(string(bs))
	}
	// Unmarshal은 err를 반환하며 json을 다시 구조체로 바꾼 값을 저장하기 위한 바이트 슬라이스가 필요하다
	var people2 []person
	// Unmarshal (JSON형식의 값과 바이트 슬라이스 구조의 값을 저장할 곳이 필요하다)
	err2 := json.Unmarshal(bs, &people2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(people2)
	fmt.Println(people2[0].First)
}
