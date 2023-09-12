package main

import "fmt"

type person struct {
	first  string
	last   string
	flaver []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	//1
	// p1 := person{
	// 	first: "Park",
	// 	last:  "steve",
	// 	flaver: []string{
	// 		"chocolate",
	// 		"vanilla",
	// 		"coke",
	// 	},
	// }

	// p2 := person{
	// 	first: "Lee",
	// 	last:  "sarah",
	// 	flaver: []string{
	// 		"cookie and cream",
	// 		"peanut",
	// 		"soda",
	// 	},
	// }

	// fmt.Println(p1.first)
	// fmt.Println(p1.last)
	// for i, v := range p1.flaver {
	// 	fmt.Println(i, v)
	// }
	// fmt.Println(p2)

	//2
	// m := map[string]person{
	// 	p1.last: p1,
	// 	p2.last: p2,
	// }
	// fmt.Println(m)

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// //3
	// t := truck{
	// 	vehicle: vehicle{
	// 		doors: 2,
	// 		color: "white",
	// 	},
	// 	fourwheels: true,
	// }

	// s := sedan{
	// 	vehicle: vehicle{
	// 		doors: 4,
	// 		color: "black",
	// 	},
	// 	luxury: true,
	// }

	// fmt.Println(t)
	// fmt.Println(s)
	// // 내부 승격
	// fmt.Println(t.doors)
	// fmt.Println(s.doors)

	//4
	// 익명 구조체
	s := struct {
		age      int
		shoes    map[string]int
		favDrink []string
	}{
		age: 25,
		shoes: map[string]int{
			"air force 1": 2,
			"dunk":        3,
			"converse":    1,
		},
		favDrink: []string{
			"cass",
			"fresh",
		},
	}

	fmt.Println(s.age)
	for k, v := range s.shoes {
		fmt.Println(k, ":", v)
	}
	fmt.Println(s.favDrink)
}
