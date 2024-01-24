package main

import (
	"fmt"
	"sync"
)

const times = 100

func main() {

	count := 0
	mutex := sync.Mutex{}

	channel := make(chan bool)

	go func() {
		for i := 0; i < times; i++ {
			mutex.Lock() //executes completely, then after it will unlock
			count++
			mutex.Unlock()
		}
		channel <- true

	}()

	go func() {
		for i := 0; i < times; i++ {
			mutex.Lock()
			count--
			mutex.Unlock()
		}
		channel <- true
	}()

	<-channel
	<-channel

	fmt.Println(count)

}
