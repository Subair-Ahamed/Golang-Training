package main

import "fmt"

type Celsius float32

func (c Celsius) ToFahrenheit() float32 { //method syntax : func (receiver with type) method_name(arguments) return type{}

	fahren := 9.0/5.0*c + 32
	return float32(fahren)

}

type Employee struct {
	Name   string
	Id     int
	Salary int
}

type Details string

func (d Details) getDetails() Employee {
	return Employee{"Raju", 1, 20000}
}

func main() {

	var c Celsius
	fmt.Println("Enter celsius:")
	fmt.Scan(&c)
	f := c.ToFahrenheit() //method name
	fmt.Println(f)

	var d Details
	get := d.getDetails()
	fmt.Println(get)
}
