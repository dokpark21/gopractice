package main

import "fmt"

func main() {
	// map = {key: value}
	// map[key type]value type
	m := map[string]int{
		"James": 32,
		"Miss":  40,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	//지정되지 않은 값읋 출력하면 zero value
	fmt.Println(m["Steve"])

	//원소 추가하기
	m["Kevin"] = 25

	for k, v := range m {
		fmt.Println(k, ":", v)
	}

	// v, ok := m["Barnabas"]
	// fmt.Println(v)  // value
	// fmt.Println(ok) // value가 존재하는지 true or false

	// Go에서 볼 수 있는 매우 자주 관용적으로 사용되는 코드 조각(매우 중요!!!!!!)
	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("There's no value in map about Barnabas!")
	}

	//map 원소 삭제하기
	delete(m, "Kevin")
	fmt.Println(m)

	// slice 출력
	xi := []int{4, 5, 6, 7, 8}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
