package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var once sync.Once

func load() {
	fmt.Println("Run the block only once")
}

func main() {

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			once.Do(load)
			//once - An object that will perform exactly one action.
			//do - Executes the function load exactly once, regardless of how many times Do is called across multiple goroutines.
		}()
	}
	wg.Wait()
}
