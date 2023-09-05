package main

import "fmt"

const (
	z     = 42 //untyped
	s int = 42 // type
)

const (
	y1 = 2020 + iota
	y2
	y3
	y4
)

func main() {
	//1
	x := 42
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	//2
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 40)
	d := (42 != 41)
	e := (42 > 33)
	f := (42 < 555)
	fmt.Println(a, b, c, d, e, f)

	//3
	fmt.Println(z, s)

	//4
	res := 42
	fmt.Printf("%d\t%b\t%#x\n", res, res, res)
	res2 := 42 << 1
	fmt.Printf("%d\t%b\t%#x\n", res2, res2, res2)

	//5
	res3 := `here is something 
	as 
	a 
	raw
	string
	"hello"
	`
	fmt.Println(res3)

	//6
	fmt.Println(y1, y2, y3, y4)
}
