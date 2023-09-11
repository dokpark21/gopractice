package main

import "fmt"

// tpye 정의 및 내부 자료 타입 선언
type person struct {
	// field : 각각의 이름과 타입 명시된 것들
	first string
	last  string
	age   int
}

// 임베디드 구조체
type secretAgent struct {
	person //익명 필드
	ltk    bool
}

func main() {
	//구조체 : 다양한 타입의 자료들을 한번에 정의 및 사용가능
	//항상 필드 이름 사용하기
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

	//익명 구조체
	//필요 없는 외부 타입이나 변수를 사용하지 않고 클린한 코드를 유지하기 위해 사용
	p3 := struct {
		first string
		last  string
	}{
		first: "Park",
		last:  "steve",
	}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(sa1)
	// 굳이 sa1.person.age로 안하고 sa1.age 가능
	// 해당 임베디드 타입의 내부 필드가 승격됨 sa1.person.age -> sa1.age
	// 승격 : inner type -> outer type(person -> secretAgent)
	fmt.Println(sa1.person, sa1.age, sa1.ltk)
	fmt.Println(p3)
}

//Go는 객체 지향 언어 인가? : go에는 타입과 메소드가 있고 객체 지향 프로그래밍 스타일을 허용하지만 타입 계층은 존재하지 않는다.
// Go는 객체 지향적 : 캡술화,  1.구조체 생성이나 상태나 필드를 가질 수 있음 2. 동작을 호용함(methods) 3. exported & unexported; viewable & not viewable
// 재사용도 가능 : 임베디드 타입에 대한 상속이 있고 다른 타입에 타입을 임베드 할 수도 있다.
// 인터페이스가 있는 다형성이 있고, 오버라이딩도 존재한다, 승격 즉 내부 승격도 존재한다

// class를 생성하지 않는다. 대신 TYPE을 생성한다.
// 인스턴스화 하지 않고 타입의 값을 생성한다. 사용자 정의 타입으로 우리는 새로운 타입을 선언할 수 있다.
