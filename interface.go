package main

import (
	"fmt"
	"math"
)

type Shape interface { //interface syntax : type interface_name interface{}
	Area() float32 //method signature
}

type Circle struct {
	radius float32
}

type Rectangle struct {
	height, width float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) Area() float32 {
	return r.height * r.width
}

func getArea(s Shape) float32 {
	return s.Area()
}
func main() {
	c := Circle{radius: 5}
	r := Rectangle{height: 10, width: 5}
	fmt.Println(getArea(c))
	fmt.Println(getArea(r))

}
