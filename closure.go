package main

import "fmt"

// This is called data isolation
// no other variable/function have access to val except counter
// val is global but only to the scope of counter()
func newCounter() func() int {
	val := 0
	return func() int {
		val++
		return val
	}
}

func main() {
	counter := newCounter()

	fmt.Println(counter()) // invoke counter
	fmt.Println(counter()) // invoke counter
	fmt.Println(counter()) // invoke counter

}
