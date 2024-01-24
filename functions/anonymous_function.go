package main

import "fmt"

func hello(msg string) { //normal function
	fmt.Println(msg)
}

func return_anonymous() func(string) { //returning an anonymous function to a function
	return func(msg string) {
		fmt.Println(msg)

	}
}

func main() {
	hello("Hello")

	//Anonymous function declaration
	func(msg string) {
		fmt.Println(msg)
	}("Anonymous function")

	//Returning anonymous function
	returning := return_anonymous()           //declaring the function to a variable
	returning("Returning anonymous function") //calling the variable as a function

}
