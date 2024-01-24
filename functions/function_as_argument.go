package main

import "fmt"

func main() {
	sqr := func(i int) int {
		return i * i
	}
	defer fmt.Print("Finished") //defer is called when the function execution is completed, it can also be written as a func.
	fmt.Println("The answer is ", double(10, sqr))
}

func double(num int, fn func(a int) int) int {
	return fn(num)
}
