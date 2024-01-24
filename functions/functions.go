package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Printf("The addition is %d\n", add(2, 3))

	fmt.Println("The maximum value is ", max(10, 8))

	m, n := div(30, 10) //func call can be written as this
	fmt.Printf("The quotient is %d and the remainder is %d\n", m, n)

	fmt.Printf("The joined string is %v\n", joinstr("This", "is", "go", "lang"))

	//func as parameter:
	sqr := func(i int) int {
		return i * i
	}
	fmt.Println("The answer is ", double(10, sqr))

	//defer: the don't exeucte until the nearby functions executed
	defer mul(3, 5)
	mul(5, 6) //this will execute 1st
	defer fmt.Println("Defer will execute after all the functions nearby executed")
}

// func declaration - 1:
func add(a, b int) int {
	sum := a + b
	return sum
}

// func declaration - 2:
var max = func(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// returning two values:
var div = func(m int, n int) (int, int) {
	return m / n, m % n
}

// variadic functions to join strings:
func joinstr(elements ...string) string { //can give unlimited parameters
	return strings.Join(elements, "/")
}

// function can be passed as a parameter:
func double(num int, fn func(int) int) int {
	return fn(num)
}

// defer:
func mul(a1, a2 int) int {
	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}
