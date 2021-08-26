package main

import (
	"fmt"
)

// will do nicely even if there is a data race

func main() {

	var data int
	go func() { data++}()
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}

//Introducing sleeps into your code can be a handy way to debug concurrent
//programs, but they are not a solution.
/*
func main() {

	go func() {
		data++
	}()
	time.Sleep(1*time.Millisecond)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
*/
