package main

import (
	"fmt"
	"sync"
)
const times = 1000

type Counter struct{
	count int
	sync.Mutex //used to lock and unlock the values received by channel
}

func main(){
	counter := Counter{count : 0}
	channel := make(chan bool) //to define an empty value by giving empty struct

    go func(){
		for i := 0; i < times; i++{
			counter.Lock()
			counter.count++ //Any code present between a call to Lock and Unlock will be executed by only one Goroutine.
			counter.Unlock()
		}
        channel <- true
	}()

	go func(){
		for i := 0;i < times; i++{
			counter.Lock()
			counter.count--
			counter.Unlock()
		}
        channel <- true
	}()

	<-channel //for increment
	<-channel //for decrement
	fmt.Println(counter.count) //will give 0
}