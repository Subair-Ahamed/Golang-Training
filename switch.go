/*---Switch-Case---*/

package main

import "fmt"

func main() {

	var day int
	fmt.Print("Enter the day:")
	fmt.Scan(&day) //dynamic input

	switch day {

	case 1, 3, 5: //Multi-case
		fmt.Println("Odd Weekday")

	case 2, 4:
		fmt.Println("Even weekday")

	case 6: //Single-case
		fmt.Println("saturday")

	case 7:
		fmt.Println("sunday")

	default:
		fmt.Println("Invalid input")
	}
}
