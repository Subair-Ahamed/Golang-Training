package main

import (
	"fmt"
	"time"
	"context"
)

func employee(ctx context.Context){ //Content interface is the core of context package

	for {
		select{
			case <- ctx.Done():        //ctx.Done() is to get the channel that is closed when the context is cancelled
			    fmt.Println("Context Cancelled...!")
                            return

	                default:
			    fmt.Println("Employee performing some task")
			    time.Sleep(1 * time.Second)
	              }
            }      
}

func main(){

	ctx,cancel := context.WithCancel(context.Background()) //context.WithCancel freturns a new context(context.Background()) and cancels the function
	defer cancel() //cancel the context

	go employee(ctx)
	time.Sleep(5 * time.Second)
}
