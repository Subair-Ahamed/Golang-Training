// Every Go program starts with the main package.
// Whenever the compiler sees the main package, it treats the program as the executable code.
package main

// we import the packages so that we can use all of its functions in our program.
import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {

	//fmt package functions:
	var inp int
	fmt.Println("enter the input:")
	fmt.Scan(&inp)
	fmt.Println("The input is ", inp)

	//string package fucntions:
	str1 := "Packages in golang"
	fmt.Println(strings.Index(str1, "i")) //gives the position of character
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToUpper(str1))

	//math package functions:
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Cbrt(8))

	//error package functions:
	err1 := errors.New("There is an error")
	if err1 != nil {
		fmt.Println(err1)
	}

}
