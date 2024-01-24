package main

import (
	"fmt"
)

func main() {

	var arr1 = [5]int{1, 3, 5, 7, 9}   //length specified
	var arr2 = []int{5, 6, 8, 3, 2, 4} //length is inferred

	for i := 0; i < len(arr1); i++ { //iterating through the array //len() to find length of array

		fmt.Printf("Index : %d  Value : %d\n", i, arr1[i])

	}

	var arr3 = [3][2]int{{2, 3}, {4, 5}, {7, 8}} //2D array
	for i := 0; i < len(arr3); i++ {
		fmt.Printf("The value at %d is %d\n", i, arr3[i])
	}

	arr1[2] = 4 //replacing the element
	fmt.Println(arr1)

	num1 := arr2[3] //assigning the array value to a variable
	num2 := arr2[5]
	fmt.Println(num1 + num2)

	//range
	for i, x := range arr2 { //x denotes the values of arr2
		fmt.Printf("%d : %d\n", i, x)
	}

	//dynamic input for array
	var length int

	fmt.Println("Enter the length of array:")
	fmt.Scan(&length)
	new_arr := make([]int, length)

	for k := 0; k < length; k++ {
		fmt.Println("Enter the values:")
		fmt.Scan(&new_arr[k])
	}

	fmt.Println("The new array is ", new_arr)
}
