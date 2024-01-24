/*---Operators in GO---*/

package main

import "fmt"

func main() {

	/*---Arithmetic Operator---*/
	a := 20
	b := 10
	ans1 := a + b
	ans2 := a - b
	ans3 := a * b
	ans4 := a / b
	fmt.Printf("The values are %d,%d,%d,%d", ans1, ans2, ans3, ans4)

	/*---Assignment Operator---*/
	var x = 50
	x += 5
	var y = 30
	y *= 2
	var z = 5
	z &= 7
	fmt.Printf("\nThe values of x,y,z are %d,%d,%d", x, y, z)

	/*---Comparision Operator---*/
	fmt.Println("\n", a == b) //prints true or false
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	/*---Logical Operator---*/
	fmt.Println(a > b && a < x)   //"&" both should be true
	fmt.Println(x >= y || x >= z) //"|" anyone condition should be true
	fmt.Println(!(x == y))        //if true, it prints false

	/*---Bitwise Operator---*/
	fmt.Println(a & b)
	fmt.Println(x | y)
	fmt.Println(x ^ z)
	fmt.Println(a >> 1)
	fmt.Println(y << 1)
}
