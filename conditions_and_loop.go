package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 15
	c := 12

	/*---Conditional Statements---*/

	if a > b {

		if a > c { //nested if

			fmt.Println("A is greater")

		} else {

			fmt.Println("C is greater")

		}

	} else if b > c {

		fmt.Println("B is greater")

	} else {

		fmt.Println("C is greater")

	}

	/*---Loops---*/

	//1.Normal method:
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//2.Similar to while loop:
	j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	//3. Break:
	for k := 0; k <= 10; k++ {
		if k == 6 {
			break //breaks the loop print the values
		}
		fmt.Println(k)
	}

	//4.Continue:
	for k := 0; k <= 10; k++ {
		if k == 6 {
			continue //continues to the loop by skipping the condition statement
		}
		fmt.Println(k)
	}

	//3.Infinite loop
	/* i := 0
	for true{
		fmt.Println(i)
		i++
	} */

}
