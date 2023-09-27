package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

// type person struct {
// 	name string
// 	age  int
// }

// func (p *person) speak() {
// 	fmt.Println("hello my name is ", p.name)
// }

// type human interface {
// 	speak()
// }

func main() {
	// //1
	// wg.Add(2)
	// go1 := func() {
	// 	fmt.Println("This is go routine 1")
	// 	wg.Done()
	// }
	// go2 := func() {
	// 	fmt.Println("This is go routine 2")
	// 	wg.Done()
	// }
	// go go1()
	// go go2()

	// wg.Wait()

	// fmt.Println("about to exit")

	// //2
	// p1 := person{
	// 	name: "Park",
	// 	age:  25,
	// }

	// saySomething := func(h human) {
	// 	h.speak()
	// }

	// saySomething(&p1)

	//3,4
	// var counter int
	// limit := 100
	// wg.Add(limit)
	// for i := 0; limit > i; i++ {
	// 	go func() {
	// 		mu.Lock()
	// 		v := counter
	// 		runtime.Gosched()
	// 		v++
	// 		counter = v
	// 		mu.Unlock()
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println("counter : ", counter)
	// fmt.Println("end of main func!!")

	//5
	// var counter int64

	// limit := 100
	// wg.Add(limit)
	// for i := 0; i < limit; i++ {
	// 	atomic.AddInt64(&counter, 1)
	// 	println("counter : ", atomic.LoadInt64(&counter))
	// 	wg.Done()
	// }
	// wg.Wait()
	// fmt.Println("end counter : ", counter)

	//6
	fmt.Println("GOOS : ", runtime.GOOS)
	fmt.Println("GOARCH : ", runtime.GOARCH)

}
