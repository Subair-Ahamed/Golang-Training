package main 

import (
	"fmt"
	"time"
	"context"
)

func employee(ctx context.Context) error{
	for{
		select{
                case <- ctx.Done():
			fmt.Println("Error occured")
			return ctx.Err() //to get the reason for context cancellation.

		default:
			fmt.Println("Operation is running..!")
			return nil
			time.After(2 * time.Second)
		      }
	   }
}

func main(){
	ctx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)   //same like WithCancel() but also takes duration
	
	defer cancel()

	err := go employee(ctx)

	if err != nil {
		fmt.Println("Main: Operation failed..!",err)
	}

}
