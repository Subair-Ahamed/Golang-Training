package main

import "fmt"

type Company struct{
	name string
	location string
	brand string
}

type Employee struct{
	ID int
	phone int
	salary int
	Company //embedding the struct inside a struct
}

func (e *Employee) setValue(){
    fmt.Println("Enter ID:")
	fmt.Scan(&e.ID)
	fmt.Println("Enter phone:")
	fmt.Scan(&e.phone)
	fmt.Println("Enter salary:")
	fmt.Scan(&e.salary)
	fmt.Println("Enter Company name:")
	fmt.Scan(&e.name)
	fmt.Println("Enter location:")
	fmt.Scan(&e.location)
	fmt.Println("Enter Company brand:")
	fmt.Scan(&e.brand)
}

func (e Employee) getValue(){
	fmt.Println("Employee ID:",e.ID)
	fmt.Println("Phone number:",e.phone)
	fmt.Println("Salary:",e.salary)
	fmt.Println("Company name:",e.name)
	fmt.Println("Location:",e.location)
	fmt.Println("Company brand:",e.brand)
    
}

func main(){
	emp := Employee{} //we can access the company struct insdie the employee struct
	emp.setValue()
	emp.getValue()

}