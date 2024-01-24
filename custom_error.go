//custom error - errors created by the user

package main

import (
	"errors"
	"fmt"
)

func Area(s float32) (error, float32) {

	if s < 0.0 {
		return errors.New("Negative value encountered"), 0 //two values returned
	}

	return nil, s * s
}

func main() {
	var s float32
	fmt.Println("Enter the side:")
	fmt.Scan(&s)

	err, area := Area(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("the area is ", area)

}
