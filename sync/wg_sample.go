package main

import (
	"fmt"
	"sync"
)

func myFunc(wg *sync.WaitGroup) {

	fmt.Println("Goroutines wwaitgroup starting......")

	wg.Done() //Decrement the counter when the goroutine completes
}

func main() {

	fmt.Println("Goroutine Waitgroup")

	var wg sync.WaitGroup

	wg.Add(1)  //it is typically called before starting a new goroutine to indicate that there is a new task to wait for. Increments the wg counter.

	go myFunc(&wg)
	go myFunc(&wg)

	wg.Wait() //it blocks until the WaitGroup counter becomes zero. It waits for all goroutines to call Done and decrement the counter to zero.

	fmt.Println("Goroutines waitgorup finished!!!!!!")
}
