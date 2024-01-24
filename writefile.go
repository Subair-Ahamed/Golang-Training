package main

import (
	"os" //package to do file operations
	"fmt"
)

func main(){

	file ,err := os.Create("samplefile.txt")  //to create a file
	content := []byte("Golang file creating, writing and reading") //content to be written in file created

	defer file.Close() //to close a file

	if err != nil{
		fmt.Println("Error creating file",err)
	} else {
		fmt.Println("File created successfully")
	}

	write , err := file.Write(content)  //to write a file
    if err != nil{
		fmt.Println("Error writing file",err)
	} else {
		fmt.Println("File written successfully",write)
	}
}