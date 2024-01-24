package main

import "fmt"

func sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

//Example-1:
func greet() func() {
	name := "raju"
	return func() {
		fmt.Println("Hi", name)
	}
}

//nested function:
func showName(input string) { //outer function

	var display = func() { //inner function
		fmt.Println("Hi", input)
	}

	display() //calling inner function
}

func main() {

	int_seq := sequence()
	fmt.Println(int_seq()) //have reference of the variable i
	fmt.Println(int_seq()) //recalls the same function

	int_seq1 := sequence()
	fmt.Println(int_seq1())
	//example-1:
	greeting := greet()
	greeting()

	//nested function:
	showName("Prakash") //calling outer function

}
