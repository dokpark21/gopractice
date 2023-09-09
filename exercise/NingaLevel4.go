package main

import "fmt"

func main() {
	//1
	// 배열 선언
	// x := [5]int{1, 2, 3, 4, 5}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)

	//2
	//slice 선언
	// x := []int{42, 43, 55, 56, 66, 43, 34, 34, 32, 43}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", x)

	//3
	// x := []int{42, 43, 44, 45, 46}
	// a := x[len(x)-1]
	// fmt.Println(x)
	// for i := 0; i < 3; i++ {
	// 	for k := 0; k < len(x); k++ {
	// 		x[k] = a + 1
	// 		a++
	// 	}
	// 	fmt.Println(x)
	// 	a = x[len(x)-1]
	// }

	//4
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x, 52)

	// fmt.Println(x)
	// x = append(x, 53, 54, 55)
	// fmt.Println(x)

	// y := []int{56, 57, 58, 59, 60}

	// x = append(x, y...)
	// fmt.Println(x)

	//5
	// x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// x = append(x[:3], x[6:]...)
	// fmt.Println(x)

	//6
	// x := make([]string, 5, 5)
	// //이렇게 하면 위에 선언한 슬라이스는 없어진다
	// x = []string{"one", "two", "three", "four", "five"}
	// fmt.Println(x)

	// //맞는 풀이
	// y := make([]string, 5, 5)
	// num := []string{"one", "two", "three", "four", "five"}
	// for i, v := range num {
	// 	y[i] = v
	// }

	// for i := 0; i < len(y); i++ {
	// 	fmt.Println(i, y[i])
	// }

	//7
	// a := []string{"James", "Bond", "Shaken, not stirred"}
	// b := []string{"Miss", "Moneypenny", "Helloooooo, James"}

	// x := [][]string{a, b}
	// fmt.Println(x)

	// for i, xs := range x {
	// 	fmt.Println("record:  ", i)
	// 	for j, v := range xs {
	// 		fmt.Println(j, v)
	// 	}
	// }

	//8,9,10
	m := map[string][]string{
		"steve_Park": []string{"ball", "TV", "PS4"},
		"kevin_choi": []string{"music", "phone", "food"},
		"sarah_jang": []string{"dog", "cat", "flower"},
	}

	m["Lebron_park"] = []string{"Lakers", "Ohio", "Good job!"}

	delete(m, "steve_Park")

	for k, v := range m {
		fmt.Print(k, ": ")
		for i, v2 := range v {
			fmt.Printf("%d:%v\t", i, v2)
		}
		fmt.Println()
	}

}
