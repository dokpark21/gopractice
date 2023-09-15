package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// speak 함수를 secretAgent 타입에 연결
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

// speak 함수를 person type에 연결
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// keyword identifier type
// 인터페이스는 값이 하나 이상의 타입의 값이 될 수 있게 한다.
// human은 secretAgent or person type이 될 수 있다(값 1개에 2개의 타입 가능!!)
// 다형성 제공
type human interface {
	// 만약 speak이라는 메소드를 가지고 있으면 human type이 된다
	speak()
}

func bar(h human) {
	// assertion (무언가를 단언함 : 아래에서는 person, secretAgent type인지에 대해 단언하고 있다.)
	switch h.(type) {
	case person:
		fmt.Println("person", h.(person).first)
	case secretAgent:
		fmt.Println("secretAgent", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	// speak 함수에 type이 secretAgent인 sa1 연결
	sa1.speak()

	p1 := person{
		first: "Dr.",
		last:  "Yse",
	}

	bar(sa1)
	bar(p1)
}
