package main //Package declaration

import "fmt" //Import Packages

func main() { //function declaration

	/* Basic Datatypes */

	var a int = 10 //"int" datatype - signed(int) & unsigned(uint)

	var b float32 = 5.25 //"float" datatype - float32 & float64

	var c string = "Hello" //"string" datatype

	var d bool = true //"bool" datatype - true or false

	var e complex64 = complex(3, 2) //"complex" datatype - complex numbers (3+2i,5+6i,...)

	fmt.Printf("The type of a is %T and value is %v", a, a) //%T to print the type of value

	fmt.Printf("\nThe type of b is %T and value is %f", b, b)

	fmt.Printf("\nThe type of c is %T and value is %s", c, c)

	fmt.Printf("\nThe type of d is %T and value is %t", d, d)

	fmt.Printf("\nThe type of e is %T and value is %v", e, e)
}
