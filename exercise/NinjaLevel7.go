package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) changeMe(ad *person, first string, last string, age int) {
	// 구조체의 주소를 참조해 필드의 값을 바꾸려면 (*value).field로 사용해야된다
	(*ad).first = first
	(*ad).last = last
	(*ad).age = age
}

func main() {
	//1
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	//2
	p1 := person{
		first: "Park",
		last:  "steve",
		age:   25,
	}

	fmt.Println(p1)

	p1.changeMe(&p1, "Jang", "james", 28)

	fmt.Println(p1)
}
