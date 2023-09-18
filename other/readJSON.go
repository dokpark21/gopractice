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
