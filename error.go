package main

import (
	"errors"
	"fmt"
)

func divide(num1, num2 int) error {

	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2) //Errorf-format printing of error
	}

	return nil
}

func main() {
	err1 := errors.New("Error Encountered")

	if err1 != nil {
		fmt.Println(err1)
	}

	err2 := divide(4, 0)

	if err2 != nil {
		fmt.Printf("error: %s", err2)

		// error not found
	} else {
		fmt.Println("Valid Division")
	}

}
