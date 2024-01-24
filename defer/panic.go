//We use the panic statement to immediately end the execution of the program.
//If our program reaches a point where it cannot be recovered due to some major errors, it's best to use panic.

package main

import "fmt"

func division(num1, num2 int) {

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func main() {
	division(4, 2)
	division(8, 0) //this will reach the panic message, hence execution will be stopped
	division(2, 8)

	fmt.Println("There is an attack")
	panic("Ending the execution!!!")

}
