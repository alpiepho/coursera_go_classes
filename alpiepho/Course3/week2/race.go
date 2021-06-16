package main

import (
	"fmt"
	"time"
)

var X int

func routine1() {
	X = 1
}

func routine2() {
	X = X + 1
}

// func testGoRoutines() {
// 	for i := 0; i < 10; i++ {
// 		go routine1()
// 		go routine2()
// 		time.Sleep(time.Second)
// 		fmt.Println(X)
// 	}
// }

func main() {
	// There is a race condition between routine1 and routine2.  If routine1 executes first, then
	// X will be 1, then routine2 will change it to 2.  If routine2 executes first, then X is 1 and
	// routine1 will also set it to 1.  So program outcome is not determinisitic.
	go routine1()
	go routine2()
	time.Sleep(time.Second)
	fmt.Println(X)
}
