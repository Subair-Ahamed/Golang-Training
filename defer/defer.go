//defer statement is used to prevent the execution of a function until all other functions execute

package main

import "fmt"

func main() {

	defer fmt.Println("1st defer statement will be executed at last")
	fmt.Println("All other statemets will execute 1st")
	defer fmt.Println("defer will be executed in a Last-in-first-out order")

	defer mul(3, 5)
	mul(5, 6) //this will execute 1st
	defer mul(4, 2)
}

//Example:
func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}
