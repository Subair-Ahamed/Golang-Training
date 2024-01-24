package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})

func produce() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Consumer:", i)
		cond.Signal()
	}
}

func consume(i int) {
	for {
		cond.L.Lock()
		cond.Wait()
		fmt.Printf("Consumer %d have done\n", i)
		cond.L.Unlock()
	}

}

func main() {

	go produce()

	for i := 1; i <= 3; i++ {
		go consume(i)
	}

	time.Sleep(7 * time.Second)
}
