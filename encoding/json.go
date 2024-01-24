package main

import (
	"fmt"
	"encoding/json"
)

func main(){

	type Employee struct{
		Name string    `json:"name"`
		Age int        `json:"age"`
		Salary int     `json:"salary"`
		Email []string `json:"email"`
	}

	entry1 := Employee{"Prakash",23,19000,[]string{"prakash@yahoo.com","prakash@gmail.com"}}

	data1,err1 := json.Marshal(entry1) //Encoding the struct which returns a byte array and error.

	if err1 != nil{
		fmt.Println("Error occured",err1)
	}
	fmt.Println(string(data1))

	entry2 := `{"Name":"Solomon","Age":40,"Salary":80000,"Email":["solomon@gmail.com","solo@hotmail.com"]}`

	data2 := &Employee{}

	err2 := json.Unmarshal([]byte(entry2),data2) //Decoding the struct by passing a byte array and the struct
	if err2 != nil{
		fmt.Println("Error occured",err2)
	}
    
	fmt.Println(data2.Name)
	fmt.Println(data2.Email)
	fmt.Println(data2.Age)
	
}