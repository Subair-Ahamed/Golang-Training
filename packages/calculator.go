package main

import (
	"packages/calculator"
	"fmt"
)

func main() {

	//we can use our own custom package - calculator functions
	fmt.Println(calculator.Add(6, 5))
	fmt.Println(calculator.Sub(6, 3))
	fmt.Println(calculator.Add(4, 2))
	fmt.Println(calculator.Add(10, 5))
}
