package main

import "fmt"

type Phone struct { //struct syntax
	Brand string
	Ram   int
	Price int
}

type Employee struct {
	ID     int
	Name   string
	Age    int
	Mobile Phone //nested struct
}

type Bank struct {
	Bname  string
	Person Employee
	AccNo  int
}

func main() {
	var ph Phone
	ph.Brand = "Samsung"
	ph.Ram = 8
	ph.Price = 25000
	fmt.Println(ph)

	var details Bank

	details = Bank{ //nested struct

		Bname: "SBI",
		Person: Employee{
			ID:   01,
			Name: "Prakash",
			Age:  23,
			Mobile: Phone{
				Brand: "Vivo",
				Ram:   16,
				Price: 40000,
			},
		},
		AccNo: 4358,
	}

	fmt.Printf("%v", details)

}
