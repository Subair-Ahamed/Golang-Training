package main

import (
	"fmt"
)

func main() {
	var name1 string

	fmt.Print("Enter name:")
	fmt.Scanf("%q", &name1)
	// m := strings.TrimSpace(name1)
	// n := strings.Fields(m)
	// x := strings.Join(n, " ")
	fmt.Println("The name is ", name1)
}
