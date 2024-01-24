package main

import "fmt"

type Employee struct {
	Name string
	ID   int
}

func main() {
	var a int = 20
	var ptr *int
	var pptr **int //double pointer

	ptr = &a //stores the address of a
	pptr = &ptr

	emp := Employee{"Prakash", 05}

	//pointer to a struct
	ptr1 := &emp

	fmt.Println(ptr1)
	fmt.Println(*ptr1)
	fmt.Println(ptr1.Name)
	fmt.Println(ptr1.ID)

	fmt.Println(ptr) //referencing - prints the address of a

	fmt.Println(*ptr) //dereferencing - prints the value of a

	fmt.Println(pptr) //prints the address of ptr

	fmt.Println(*pptr) //prints the address of value stored in ptr

	fmt.Println(**pptr) //prints the value of ptr

}
