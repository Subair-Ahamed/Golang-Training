package main

import (
	"fmt"
	"sort"
)

func main() {

	arr1 := [6]string{"This", "is", "go", "lang", "training", "session"}
	arr2 := [5]int{20, 11, 3, 6, 35}

	//slice declaration
	var slice1 = arr1[0:4]
	slice2 := arr2[:3] //20,11,3
	slice3 := arr1[2:]
	slice4 := arr2[:]   //all the elements from arr2
	slice5 := arr2[3:5] //6,35

	fmt.Println("Slice 1 :", slice1)
	fmt.Println("Slice 2 :", slice2)
	fmt.Println("Slice 3 :", slice3)
	fmt.Println("Slice 4 :", slice4)

	fmt.Printf("Len of slice3 : %d\n", len(slice1))      //length of the slice
	fmt.Printf("Capacity of slice3 : %d\n", cap(slice1)) //capacity of the size

	new_slice1 := append(slice2, 12, 13) //adding elements to the end of slice
	fmt.Println("New slice is ", new_slice1)

	copy(slice5, slice2) //copy(copyto , copyfrom) - copying elements from one slice to another until reeaches the capacity
	fmt.Println("The copied slice is ", slice5)

	//range:
	for _, x := range slice1 { //iterating through the slice
		fmt.Printf("The elements in slice1 are %v\n", x)
	}

	sort.Ints(slice2) //sort keyword is used to sort a slioce in ascending order //Ints() is used for integers & Strings() for string
	fmt.Println(slice2)
}
