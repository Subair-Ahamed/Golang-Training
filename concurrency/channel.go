//Channels - are link between goroutines and used for synchronizing
package main

import "fmt"

func main(){

    channel := make(chan int) //to create a chasnnel
	go counter(channel)

	inp := <-channel //inp variable will receive an input from the channel
	fmt.Println(inp)

	inp = <-channel
	fmt.Println(inp)

	inp = <-channel
	fmt.Println(inp)
}

func counter(name chan int){ //channel passed as argument
	for i := 0; i<5 ;i++{
		name <- i //values received by the channel
	}
}