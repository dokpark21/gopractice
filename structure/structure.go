package main

import "fmt"

// tpye 정의 및 내부 자료 타입 선언
type person struct {
	first string
	last  string
	age   int
}

// 임베디드 구조체
type secretAgent struct {
	person
	ltk bool
}

func main() {
	//구조체 : 다양한 타입의 자료들을 한번에 정의 및 사용가능
	p1 := person{
		first: "Jamses",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   35,
	}

	sa1 := secretAgent{
		person: person{
			//inner type(승격 : inner type -> outer type)
			first: "Park",
			last:  "steve",
			age:   25,
		},
		ltk: true,
	}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(sa1)
	// 굳이 sa1.person.age로 안하고 sa1.age 가능
	// 해당 임베디드 타입의 내부 필드가 승격됨 sa1.person.age -> sa1.age
	// 승격 : inner type -> outer type(person -> secretAgent)
	fmt.Println(sa1.person, sa1.age, sa1.ltk)
}
