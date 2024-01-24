/*----Declaration of values----*/

package main

import "fmt"

func main() {

	/*---SYNTAX: var variablename type = value---*/

	var num1 int = 10 //"var" keyword is used to declare variables with its datatype

	var str1 string = "Hello" //type is string

	var flt1 float32 = 3.16 //type is float

	var comp1 complex128 = complex(5, 6) //type is complex

	const PI = 3.14 //"const" keyword is used to declare constant values which cannot be changed

	num2 := 30 //":=" type is inferred -  compiler decides the type of variable

	/*---Declaring Multiple values---*/

	var x, y, z = 1, "Go", true

	var (
		a int     = 50
		b float32 = 4.5
		c string  = "good"
	)

	fmt.Printf("Type of num1 is %T , num2 is %T", num1, num2) //Printf is used for formatted printing

	fmt.Printf("\nType of str1 is %T , flt1 is %T  and comp1 is %T", str1, flt1, comp1)

	fmt.Printf("\nType of x,y,z is %T, %T, %T", x, y, z)

	fmt.Printf("\nValues of a,b,c are %v, %f, %s", a, b, c)

	fmt.Println("\nAdd of 2 numbers are ", num1+num2) //Println is used to print the direct value

	fmt.Println("The value of PI is ", PI)

}
