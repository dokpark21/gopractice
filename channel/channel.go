package main

import "fmt"

func main() {
	// // wrong case
	// // channel 생성
	// c := make(chan int)

	// // channel에 42를 할당
	// c <- 42
	// // 이 부분에서 채널이 막힌다(수신자가 설정되지 않았기 떄문에 여기서 막힌다)
	// // channel에서 값을 가져와 프린트
	// fmt.Println(<-c)
	// // 위 코드를 실행하면 모든 go routine이 슬립 상태라 교착 상태에 빠졌다고 하는데 이유는 채널이 막혔기 때문이다

	// correct case
	// 채널은 go routine으로 값을 주고 받는다.
	// 채널 생성과 아래의 go routine은 동시에 실행된다.
	c := make(chan int)

	// 아래의 코드는 채널의 값을 넣을 준비가 되어 있고 그 아래의 코드는 채널의 값을 가져올 준비가 되어 있다. 이 코드들은 동시에 실행된다.
	go func() {
		c <- 42
	}()
	// 값을 단순히 출력하는게 아니라 c에서 값을 빼낸다.
	fmt.Println(<-c)
	// 한번 빼냈기 때문에 다시 출력하면 값은 없다.
	// fmt.Println(<-c)

	// buffer channel : 채널의 값을 수신할 대상이 없어도 값이 채널안에 머물 수 있게 해준다.
	d := make(chan int, 1) // 값이 한개 머물 수 있게 해준다.

	d <- 10

	fmt.Println(<-d)

	d <- 15

	fmt.Println(<-d)
	fmt.Printf("%T\n", d)
	// directional channel(방향성이 존재하는 채널 === 단방향 채널)
	// 채널은 송신전용이나 수신전용으 설정할 수 있다.(단방향 채널)
	cs := make(chan<- int) // 송신전용 채널
	cr := make(<-chan int) // 수신전용 채널

	// 방향이 있는 채널을 양방향 채널에 할당은 불가능(구체화된 타입을 일반 타입에 할당은 불가능)
	// 하지만 양방향 채널을 방향이 있는 채널에 할당은 가능(일반 타입을 구체화된 타입에 할당 가능)
	// fmt.Printf("c\t%T\n", (chan int)(cr))
	// fmt.Printf("c\t%T\n", (chan int)(cs))

	cs = c
	cr = c
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

	// 방향성 채널 사용

}

// send
func foo() {

}

// receive
func bar() {

}
