package main

import (
	"encoding/json"
	"fmt"
)

// byte 슬라이스 <-> 문자열
var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

type animal struct {
	Name  string
	Order string
}

var animals []animal

func main() {
	// jsonBlob의 값을 animals, 즉 Animal 슬라이스로 언마샬링(에러를 반환하고 바이트 슬라이스와 슬라이스의 주소를 받아 언마샬링)
	err := json.Unmarshal(jsonBlob, &animals)
	// err가 있다면 프린트
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

// 프로그램에 무언가 들어오고 나가자마자 인코딩과 디코딩을 이용해서 그것을 JSON으로 바로 전송 가능하다. Marshal이나 Unmarshal처럼 변수에 할당할 필요는 없다.
