package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	fmt.Println(s)
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte("dsad"))

	if err != nil {
		fmt.Println(err)
	}

	var newBs []person

	err = json.Unmarshal(bs, &newBs)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("newBs : ", newBs)
	}
}
