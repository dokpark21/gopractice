package main

import (
	"errors"
	"fmt"
	"log"
)

// type person struct {
// 	First   string
// 	Last    string
// 	Sayings []string
// }

// type customErr struct {
// 	info string
// }

type sqrtError struct {
	lat  string
	long string
	err  error
}

func main() {
	//1
	// p1 := person{
	// 	First:   "James",
	// 	Last:    "Bond",
	// 	Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	// }

	// bs, err := json.Marshal(p1)
	// if err != nil {
	// 	// fmt.Println(err)
	// 	// 이 프로그램은 데이터를 JSON으로 변환하지 못하면 필요가 없는 프로그램이기 떄문에 변환에 실패한다면 바로 프로그램을 종료해야 한다
	// 	// 단순히 에러를 출력하기 보다는 바로 프로그램을 종료하는 편이 좋다
	// 	// 항상 프로그램의 로직을 살핀다음 그에 맞는 에러 처리를 해줘야한다.
	// 	log.Fatalln("JSON did not marshal - here's the error : ", err)
	// }
	// fmt.Println(string(bs))

	//2
	// p1 := person{
	// 	First:   "James",
	// 	Last:    "Bond",
	// 	Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	// }

	// bs, err := toJSON(p1)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(bs))

	//3
	// c1 := customErr{
	// 	info: "need more a coffee",
	// }

	// foo(c1)

	//4
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

//2
// func toJSON(a interface{}) ([]byte, error) {

// 	bs, err := json.Marshal(a)
// 	if err != nil {
// 		var newError = fmt.Errorf("heres new error : %v", err)
// 		return []byte{}, newError
// 	}
// 	return bs, nil
// }

// 3
// func (ce customErr) Error() string {
// 	return fmt.Sprintf("here is the error: %v", ce.info)
// }

// func foo(e error) {
// 	fmt.Println("foo ran - ", e, "\n")
// }

// 4
func (se sqrtError) Error() string {
	return fmt.Sprintf("lat: %v, long: %v, err: %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, sqrtError{
			lat:  "52,13",
			long: "42.132",
			err:  errors.New("here is the sqrt error!"),
		}
	}
	return 42, nil
}
