package main

import (
	"fmt"
	"strings"
)

// variadic function can take infinite arguments
func mul(nums ...int) int {

	result := 1
	for _, x := range nums {
		result *= x
	}
	return result //return is must

}

func joinstr(elements ...string) string { //function to join strings
	return strings.Join(elements, "-")
}

func main() {

	fmt.Println("The ans is ", mul(2, 4, 3, 5, 7, 6))

	//variadic functions can also be applied to slices
	nums := []int{3, 2, 4, 5}
	fmt.Println("The ans is ", mul(nums...)) //syntax to give slice as an argument

	//join strings:
	fmt.Printf("The joined string is %v", joinstr("these", "are", "called", "variadic", "functions"))

}
