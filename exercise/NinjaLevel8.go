package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// 1,2
// json으로 encoding 하기 위해서는 반드시 필드명 맨 앞글자가 대문자
type user struct {
	First string
	Age   int
}

type person struct {
	Name string
	Age  int
}

type ByAge []person

type ByName []person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	//1,2
	fmt.Println("1,2.")
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Monneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

	var newBs []user

	err = json.Unmarshal(bs, &newBs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("newBs : ", newBs)
	}

	//3
	fmt.Println("\n\n3.")
	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

	//4
	fmt.Println("\n\n4.")
	people := []person{
		{"Park", 25},
		{"Kim", 24},
		{"Jung", 27},
	}

	sort.Sort(ByAge(people))

	fmt.Println("Sort by age : ", people)

	sort.Sort(ByName(people))

	fmt.Println("Sort by name : ", people)

}
