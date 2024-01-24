package main

import (
	"os" //package to do file operations
	"fmt"
)

func main(){

	file ,err := os.Open("samplefile.txt")  //to open a file
	if err != nil{
		fmt.Println("Error opening file",err)
	} else {
		fmt.Println("File opened successfully")
	}

	defer file.Close() //to close a file(defer executes at last)

	content := make([]byte, 500) //variable to read the contents of the file as a byte array with 500 capacity.

	read , err := file.Read(content)  //to read a file

    if err != nil{
		fmt.Println("Error reading file",err)
	} else {
		fmt.Println("success : ",string(content))
	}
	fmt.Println("File read successfully",read)
}