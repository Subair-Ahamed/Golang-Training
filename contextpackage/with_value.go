package main

import (
	"fmt"
	"context"
	"time"
)

type key int

const userIDKey key = iota //iota is a pre-declared identifier in Go that is used in const declarations to simplify the definition of incrementing values.

func operation(ctx context.Context){

	//ctx.Value() is to retrieve the value associated with a specific key from a Context.
	UserID, ok := ctx.Value(userIDKey).(int) //(int) is a type assertion. It attempts to assert that the value obtained from the context is of type 'int'.

	if !ok{
		fmt.Println("Error occured : UserdID not found")
		return
	}

	fmt.Println("Processing request for userID",UserID)
}

func main(){

	ctx := context.WithValue(context.Background(),userIDKey,453) //returns a new context , key, value

	go operation(ctx)

	time.Sleep(2 * time.Second)

}