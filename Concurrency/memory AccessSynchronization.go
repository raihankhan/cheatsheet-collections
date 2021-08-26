package main

import (
	"fmt"
	"sync"
)
// synchronize access to the memory between your critical sections
// we have solved our data race, we haven’t actually solved our race condition
func main() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}