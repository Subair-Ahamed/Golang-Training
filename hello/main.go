package main

import "fmt"

func main() {
	var a byte
	fmt.Printf("Enter a:")
	fmt.Scanf("%d", a)

	var b byte
	fmt.Printf("Enter b:")
	fmt.Scanf("%d", b)
	var c = a + b
	fmt.Println(c)
}
