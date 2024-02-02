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
	radius float32
}

type Rectangle struct {
	height, width float32
}

//Method - Use CamelCase with the first letter capitalized
func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float32 {
	return r.height * r.width
}

//Function - names should use camelCase, where the first letter of each word after the first is capitalized.
func getArea(s Shape) float32 {
	return s.Area()
}

func main() {
	//Variable - Variable names should use camelCase, which means the first letter of each word is capitalized except for the first word. 
	c := Circle{radius: 5}
	r := Rectangle{height: 10, width: 5}
	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}
