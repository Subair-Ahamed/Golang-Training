package main

import (
	"fmt"
	"sync"
	"time"
)
var(
value = 0
rwmutex  sync.RWMutex
)

func readValue(id int) {
	rwmutex.RLock() //allows multiple readers but prevents writers
	defer rwmutex.RUnlock()

	fmt.Printf("Reader %d: Value is %d\n", id, value)
}

func writeValue(id int) {
	rwmutex.Lock()
	defer rwmutex.Unlock()

	value++
	fmt.Printf("Writer %d: Incremented value to %d\n", id, value)
}

func main() {


	// Run multiple readers and writers
	for i := 0; i < 5; i++ {
		go readValue(i)
		go writeValue(i)
	}

	time.Sleep(2 * time.Second) // Wait for goroutines to finish
}
