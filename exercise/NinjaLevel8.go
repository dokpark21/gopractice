package main

import (
	"encoding/json"
	"fmt"
)

// 1,2
// json으로 encoding 하기 위해서는 반드시 필드명 맨 앞글자가 대문자
type user struct {
	First string
	Age   int
}

func main() {
	//1,2
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
}
