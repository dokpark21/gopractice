package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

// 메소드 타입은 타입이 구현하는 인터페이스를 결정한다.(매우 중요!!)
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	// info는 shape를 받기 떄문에 circle이 shape를 구현하지 않는 이상 사용 불가능(메소드가 포인터 리시버를 가지고 있기 떄문)
	// 만약 method area가 포인터 리시버를 가지고 있지 않다면 circle은 shape를 구현하는 것이 맞기 떄문에 info 실행 가능
	// info(c)

	fmt.Println(c.area())
}
