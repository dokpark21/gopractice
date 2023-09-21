package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// sort.Sort에는 3가지 함수가 존재한다.
// 1. Len() slice의 길이를 반환한다
// 2. Swap() 원소들끼리의 교체를 담당한다.
// 3. Less() 원소들끼리 비교 후 작은 값이 작은지에 대한 여부를 반환한다.

func main() {
	xi := []int{4, 7, 3, 42, 88, 14}
	xs := []string{"James", "Q", "M", "MonneyPenny"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println("after sort: ", xi)
	fmt.Println("after sort xs : ", xs)

	people := []Person{
		{"steve", 54},
		{"Bob", 26},
		{"Kevin", 52},
		{"James", 43},
	}
	fmt.Println("Before sort : ", people)
	sort.Sort(ByAge(people))
	fmt.Println("After sort : ", people)

	// 어떤 것을 기준으로 sort 하기

}
