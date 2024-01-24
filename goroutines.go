//Goroutines are used for concurrent and asynchronous programming
package main

import (
	"fmt"
	"time"
)

func display(str string){
	for i:=1;i<=3;i++{
        fmt.Printf("%v : %v\n",str,i)
		time.Sleep(1 * time.Second)
	}
}

func main(){
	go display("Virat") //moves to next statement before execution of current statement
	go display("Dhoni")
	go display("Rohit")
	time.Sleep(2 * time.Second)
	fmt.Println("Executed")
}