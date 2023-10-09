package main

import (
	"fmt"
	"log"
)

// var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	v, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(v, err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// Errorf는 errors.New를 반환한다.(format print 형식으로!)
		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	return 42, nil
}
