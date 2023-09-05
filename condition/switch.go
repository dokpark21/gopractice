package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		// fallthrough //go는 다른 언어들 처럼 자동으로 쭉 내려가고 break로 제어 하는 것이 아닌 참인 것이 존재한다면 멈춤, fallthrough를 사용하면 아래것을 실행함
	default:
		fmt.Println("This is default case") //fallthrough가 없거나 위의 모든 조건이 false일 때 출력 됨
	}

	//변수나 여러개의 case를 사용하는 것도 가능
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money")
	case "Bond2":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}
}
