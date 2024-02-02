package main

import (
	"fmt"
	"math"
)

//Interface - names should represent the functionality that the interface defines. Use nouns or noun phrases.
type Shape interface { 
	Area() float32 //method signature
}

type Circle struct { //interface syntax : type interface_name interface{method signature}
	Radius float32
}

type Rectangle struct {
	//When we give fields in lowercase in a struct like (height, width), it becomes unexported, which can be accessed only within this package.
	//When we give fields in uppercase in a struct like (Height, Width), it becomes exported, which can be accessed from other packages.
 	Height, Width float32 
}

//Method - Use CamelCase with the first letter capitalized
func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

//Function - names should use camelCase, where the first letter of each word after the first is capitalized.
func getArea(s Shape) float32 {
	return s.Area()
}

func main() {
	//Variable - Variable names should use camelCase, which means the first letter of each word is capitalized except for the first word. 
	c := Circle{Radius: 5}
	r := Rectangle{Height: 10, Width: 5}
	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}
